package main

import (
	"os"
	"bufio"
	"strings"
	"sync"
	"fmt"
	"time"
	"context"
)

type words struct {
	sync.Mutex
	found map[string]int
}

func NewWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, w *words) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		w.add(word, 1)
	}
	return scanner.Err()
}

func concurrency(guard *sync.WaitGroup) {
	defer guard.Done()
	fileNames := []string{"main9.go", "main8.go"}

	var wg sync.WaitGroup
	w := NewWords()

	for _, f := range fileNames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := tallyWords(f, w); err != nil {
				fmt.Println(err)
			}
		}(f)
	}
	wg.Wait()

	fmt.Println("words that appear more than once")
	w.Lock()
	defer w.Unlock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d \n", word, count)
		}
	}
}

func channel(guard *sync.WaitGroup) {
	defer guard.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(ctx, echo)
	for {
		select {
		case data := <-echo:
			os.Stdout.Write(data)
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		}
	}
}
func readStdin(ctx context.Context, out chan<- []byte) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel from out side")
			return
		default:
			data := make([]byte, 1024)
			l, _ := os.Stdin.Read(data)
			if l > 0 {
				out <- data[0:l]
			}
		}
	}
}
func main() {
	guard := new(sync.WaitGroup)
	guard.Add(1)
	go concurrency(guard)

	guard.Add(1)
	go channel(guard)

	guard.Wait()
	time.Sleep(2 * time.Second)
}
