package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var Wait sync.WaitGroup
var Counter int64 = 0

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		// Counter = Counter + 1  // 会出现race  竞争
		fmt.Println(atomic.AddInt64(&Counter, 1)) // atomic  原子操作  是轻量级的锁
		time.Sleep(1 * time.Nanosecond)
	}

	Wait.Done()
}
