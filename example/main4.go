package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan interface{})
	var c2 chan interface{}
	var ii interface{}
	go func() {
		c2 = c
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			// ii = i
			select {
			case c2 <- ii:
				c2 = nil
			}
		}
		close(c2)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
