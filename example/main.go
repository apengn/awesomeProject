package main

import (
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/uber/makisu/lib/log"
)

func main() {
	go func() {
		ch := make(chan string, 0)

		go func() {

			fmt.Println("pp")
			time.Sleep(2 * time.Second)

			ch <- "wwwww"
			// close(ch)
		}()

		go func() {
			for e := range ch {
				if e == "" {
					fmt.Println("ch nil")
				}
				fmt.Println("====", e)
			}

			fmt.Println("end")
		}()
	}()

	log.Error(http.ListenAndServe(":6060", nil))
}
