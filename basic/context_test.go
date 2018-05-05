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
	d := time.Now().Add(50 * time.Millisecond)

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
