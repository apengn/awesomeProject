package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func TestCond(t *testing.T) {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	go func() {
		cond.L.Lock()
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			if i == 3 {
				time.Sleep(2 * time.Second)
				cond.Wait()
			}
			fmt.Println(i)
		}
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		for i := 20; i < 30; i++ {
			time.Sleep(1 * time.Second)
			if i == 23 {
				time.Sleep(2 * time.Second)
				cond.Wait()
			}
			fmt.Println(i)
		}
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		for i := 10; i < 20; i++ {
			time.Sleep(1 * time.Second)
			if i == 12 {
				time.Sleep(2 * time.Second)
				cond.Wait()
			}
			fmt.Println(i)
		}
		cond.L.Unlock()
	}()

	go func() {
		time.Sleep(20 * time.Second)
		cond.Broadcast()
		fmt.Println("Signal")
	}()

	time.Sleep(100 * time.Second)
	fmt.Println("wait")
}

func TestChannel(t *testing.T) {
	ch := make(chan interface{}, 1)

	go func() {
		time.Sleep(4 * time.Second)
		//ch <- 1
		//ch <- 2
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("done")
	fmt.Println(len(ch))
	//fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func TestSilce(t *testing.T) {
	var slice []int
	for e := range slice {
		pretty.Log(e)
	}
	pretty.Println(len(slice))
}
