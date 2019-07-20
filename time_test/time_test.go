package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	now := time.Now()

	time.Sleep(time.Second)
	fmt.Printf("%f", time.Now().Sub(now).Seconds())
}

func TestXorm(t *testing.T) {

}

func TestTick(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	c := ticker.C
	for now := range c {
		fmt.Printf("%v %s\n", now, "")
		ticker.Stop()
		break
	}
}
