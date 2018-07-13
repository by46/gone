// +build go1.8

package basic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("done")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestDeadlineContext(t *testing.T) {
	d := time.Now().Add(5000 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestTimeoutContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestValueContext(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	key := favContextKey("language")
	ctx := context.WithValue(context.WithValue(context.Background(), key, "Go"), favContextKey("color"), "red")
	f(ctx, key)
	f(ctx, favContextKey("color"))
}

func Monitor(ctx context.Context, worker int) {
	go func(ctx context.Context, worker int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("end monitor ", worker)
				return
			default:
				fmt.Println("monitoring ", worker)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx, worker*10+worker)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end monitor ", worker)
			return
		default:
			fmt.Println("monitoring ", worker)
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go Monitor(ctx, 1)
	go Monitor(ctx, 2)
	go Monitor(ctx, 3)

	time.Sleep(5 * time.Second)
	fmt.Println("stop monitoring")
	cancel()
	time.Sleep(time.Second)
}

func TestTimeoutContext2(t *testing.T) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("end")
					return
				default:
					fmt.Println("monitor")
					time.Sleep(time.Millisecond * 200)
				}
			}
		}(ctx)
		fmt.Println("waiting")
		time.Sleep(time.Millisecond * 1500)
	}()

	time.Sleep(time.Second * 2)
}

func TestMultipleCancelContext(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	ctx2, cancel := context.WithCancel(ctx)

	defer cancel()
	fmt.Println(ctx2)
}

type ValueContext struct {
	context.Context
	key, value interface{}
}

func (c *ValueContext) String() string {
	return fmt.Sprintf("%v.ValueContext(%#v, %#v)", c.Context, c.key, c.value)
}

func (c *ValueContext) Value(key interface{}) interface{} {
	if c.key == key {
		return c.value
	}
	return c.Context.Value(key)
}

func TestCustomerContext(t *testing.T) {

	ctx3, _ := context.WithCancel(context.Background())
	ctx := &ValueContext{Context: ctx3, key: "benjamin", value: "red"}
	ctx2, cancel := context.WithCancel(ctx)

	defer cancel()
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("benjamin"))
	time.Sleep(time.Second * 2)
	cancel()
}

func Process(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	go CPUProcess(ctx)
	go MemoryProcess(ctx)
	go DiskProcess(ctx)

}

func CPUProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cpu end")
			return
		default:
			fmt.Println("process cpu")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func MemoryProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("memory end")
			return
		default:
			fmt.Println("process memory")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func DiskProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("disk end")
			return
		default:
			fmt.Println("process disk")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func TestMultipleContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go Process(ctx)

	select {
	case <-ctx.Done():
		fmt.Printf("done")
	case <-time.NewTimer(time.Second * 3).C:
		fmt.Printf("sleep 3 time")
	}
	cancel()
}

type message struct {
	responseChan chan<- int
	parameter    string
	ctx          context.Context
}

func ProcessMessages(work <-chan message) {
	for job := range work {
		select {
		case <-job.ctx.Done():
			continue
		default:
		}
		result := len(job.parameter)
		select {
		case <-job.ctx.Done():
		case job.responseChan <- result:
		}
	}
}

func NewRequest(ctx context.Context, input string, q chan<- message) {
	r := make(chan int)
	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q cloud see message")
		return
	case q <- message{
		responseChan: r,
		parameter:    input,
		ctx:          ctx,
	}:
	}

	select {
	case out := <-r:
		fmt.Printf("the len of %s is %d\n", input, out)
	case <-ctx.Done():
		fmt.Println("Context ended before q cloud process message")
	}
}

func TestChannelProcessModel(t *testing.T) {
	q := make(chan message)
	go ProcessMessages(q)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Nanosecond)
	defer cancel()
	time.Sleep(time.Microsecond)
	
	NewRequest(ctx, "hei", q)
	NewRequest(ctx, "hello world", q)
	close(q)
}
