package test

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

// https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html
// https://www.ardanlabs.com/blog/2018/12/goroutine-leaks-the-abandoned-receivers.html

// goroutine leaks
func TestGoroutineLeak(t *testing.T) {
	fmt.Println("Hello")
	go func() {
		for {
			fmt.Println("Goodbye")
		}
	}()
}

func TestWg(t *testing.T) {

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		wg.Done()
	}()
	shutdown(wg, ctx)

}

func shutdown(wg *sync.WaitGroup, ctx context.Context) {

	ch := make(chan interface{})

	go func() {
		wg.Wait()
		close(ch) // only read
	}()

	select {
	case <-ch:
		fmt.Println("execute done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

func TestGoroutineLeaks(t *testing.T) {
	records := make([]string, 0)
	for i := 0; i < 8; i++ {
		records = append(records, string(i))
	}
	processRecords(records)
}

// processRecords is given a slice of values such as lines
// from a file. The order of these values is not important
// so the function can start multiple workers to perform some
// processing on each record then feed the results back.
func processRecords(records []string) {

	// Load all of the records into the input channel. It is
	// buffered with just enough capacity to hold all of the
	// records so it will not block.

	total := len(records)
	input := make(chan string, total)
	for _, record := range records {
		input <- record
	}
	//close(input)
	fmt.Println("input====", len(input))
	//close(input) // What if we forget to close the channel?

	// Start a pool of workers to process input and send
	// results to output. Base the size of the worker pool on
	// the number of logical CPUs available.

	output := make(chan string, total)
	workers := runtime.NumCPU()
	fmt.Println("NumCPU====", workers)
	for i := 0; i < workers; i++ {
		go worker(i, input, output)
	}

	// Receive from output the expected number of times. If 10
	// records went in then 10 will come out.

	for i := 0; i < total; i++ {
		result := <-output
		fmt.Printf("[result  ]: output %s\n", result)
	}
}

// worker is the work the program wants to do concurrently.
// This is a blog post so all the workers do is capitalize a
// string but imagine they are doing something important.
//
// Each goroutine can't know how many records it will get so
// it must use the range keyword to receive in a loop.
func worker(id int, input <-chan string, output chan<- string) {
	for v := range input {
		fmt.Printf("[worker %d]: input %s\n", id, v)
		output <- strings.ToUpper(v)
	}
	fmt.Printf("[worker %d]: shutting down\n", id) // 不会执行
}
