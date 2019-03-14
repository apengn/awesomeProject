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
