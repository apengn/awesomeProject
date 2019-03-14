package test

import (
	"fmt"
	"sync"
	"testing"
)

type Persion struct {
	Name string
}

// 如果put  中取完了   就重新new一个新的
func TestPool(t *testing.T) {
	pool := sync.Pool{New: func() interface{} {
		return &Persion{Name: "test"}
	}}

	i := Persion{}
	i.Name = "wp"

	pool.Put(&i)
	get := pool.Get()

	persion, ok := get.(*Persion)
	t.Log(ok)
	if ok {
		t.Log(persion.Name)
	}

	get = pool.Get()
	persion, ok = get.(*Persion)
	t.Log(ok)
	if ok {
		t.Log(persion.Name)
	}
}

func TestName(t *testing.T) {
	t.Log(test())
}

func test() string {
	s := "a"

	for i := 0; i < 4; i++ {
		if i == 9 {
			return s
		}

	}
	s = "ff"
	return s
}

func TestNil(t *testing.T) {
	Nil([]int{})
}

func Nil(selice []int) {
	for k := range selice {
		fmt.Println(k)
	}
	fmt.Println(selice)
}
