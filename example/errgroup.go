package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

func main1x() {

	group := errgroup.Group{}
	group.Go(func() error {

		return nil
	})

	group.Go(func() error {

		return nil
	})

	if err := group.Wait(); err != nil {

	}
}

func main() {
	xx()

}

func fff() {

	group := singleflight.Group{}

	doChan := make(<-chan singleflight.Result, 1)
	go func() {
		for {
			select {
			case c := <-doChan:
				fmt.Println(c)
			}
		}
	}()
	doChan = group.DoChan("ww", func() (i interface{}, e error) {

		return "wwwww", nil
	})
	doChan = group.DoChan("ww", func() (i interface{}, e error) {

		return "wwwww", nil
	})
	doChan = group.DoChan("ww", func() (i interface{}, e error) {

		return "wwwww", nil
	})

	time.Sleep(100 * time.Second)
}
func xx() {

	var requestGroup singleflight.Group
	var ii int = 0

	f := func() (interface{}, error) {

		return func() (interface{}, error) {
			n := rand.Int63n(1000)
			return n, nil
		}()
	}
	for ; ii < 100; ii++ {
		go func() {
			v, err, shared := requestGroup.Do("kkkkk", f)
			fmt.Println(v, err, shared)
		}()
	}

	requestGroup.Forget("kkkk")
	time.Sleep(10 * time.Second)

}
