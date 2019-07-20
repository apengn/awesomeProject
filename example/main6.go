package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for true {
		duration := time.Duration(rand.Int63n(int64(time.Second)))
		duration = time.Second*2 + duration
		fmt.Println(duration)
		time.Sleep(duration)
	}
}
