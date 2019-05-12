package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"gitee.com/wisecloud/wise-resource-manager/pkg/domain"
)

type Zoom struct {
	// Name string
	// Id   string
}

type I interface {
	GetName() string
	GetName2() string
}
type Zoom2 struct {
}

type Dog struct {
	Zoom
	Zoom2
}
type Cat struct {
	Zoom
}

func (z *Zoom) GetName() string {
	return "zoom"
}

func (z *Zoom2) GetName2() string {
	return "zoom2"
}

func TestDog(t *testing.T) {
	dog := Dog{}
	// dog.Name = "xxx"
	// dog.Id = "33"
	b, _ := json.Marshal(dog)
	t.Log(string(b))
	var i I
	i = &dog
	t.Log(i.GetName2())

	t.Run("test", func(t *testing.T) {
		alwaysFalse := func() bool {
			return false
		}
		switch alwaysFalse(); false {
		case true:
			println("真")
		case false:
			println("假")
		}
	})

	t.Run("ss", func(t *testing.T) {
		gen := func(ctx context.Context) <-chan int {
			dst := make(chan int)
			n := 1
			go func() {
				for {
					select {
					case <-ctx.Done():
						t.Log("done")
						return // returning not to leak the goroutine
					case dst <- n:
						n++
					}
				}
			}()
			return dst
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // cancel when we are finished consuming integers

		for n := range gen(ctx) {
			fmt.Println(n)
			if n == 5 {
				break
			}
		}

		// time.Sleep(100*)
	})
}

func ValidateDuplicate(groups []*domain.IngressGroup, v string) bool {
	for _, group := range groups {
		if v == group.GroupName {
			return true
		}
	}
	return false
}
func TestValidateDuplicate(t *testing.T) {
	vs := []string{"redis", "mysql", "redis", "mysql"}
	ig := []*domain.IngressGroup{}
	for _, v := range vs {
		if !ValidateDuplicate(ig, v) {
			ig = append(ig, &domain.IngressGroup{GroupName: v})
		}
	}

	for _, v := range ig {
		t.Log(v.GroupName)
	}

	s := []*struct {
		Name string
		Age  int
	}{{Name: "name1", Age: 4}, {Name: "name2", Age: 6}}

	t.Run("test", func(t *testing.T) {

		for _, e := range s {
			if e.Age == 4 {
				e.Name = "www"
			}
			if e.Age == 6 {
				e.Name = "pppp"
			}
		}

		b, _ := json.Marshal(s)
		t.Log(string(b))

	})
}
