package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func TestEnv(t *testing.T) {

	t.Run("env", func(t *testing.T) {
		os.Clearenv()
		if err := os.Setenv("ENV", "111"); err != nil {
			t.Fatal(err)
		}

		environ := os.Environ()
		syscall.Unsetenv("ENV")
		pretty.Print(environ)
	})

	t.Run("splicN", func(t *testing.T) {
		v := "44444-43333-4333s-fsdf"
		t.Log(strings.SplitN(v, "-", 1))
	})

	t.Run("nil", func(t *testing.T) {
		var s []string

		if len(s) == 0 {
			t.Log("fff")
		}
	})

	t.Run("channel", func(t *testing.T) {

		go func() {
			ch := make(chan string, 1)

			go func() {

				fmt.Println("pp")
				time.Sleep(2 * time.Second)

				// ch <- "wwwww"
				// close(ch)
			}()

			go func() {
				for e := range ch {
					if e == "" {
						t.Log("ch nil")
					}
					t.Log("====", e)
				}
			}()
		}()

		time.Sleep(5 * time.Second)

	})
}

type Interface interface {
	Test() string
}

type Test1 struct {
	V string
}

func (t *Test1) Test() string {
	return t.V
}

func T(p Interface) {
	switch p.(type) {
	case *Test1:
		s := p.(*Test1)
		s.V = "hello"
	}
}

func Test_InterFace(t *testing.T) {
	test1 := &Test1{}
	T(test1)
	t.Log(test1.V)
}

func Test_xxx(t *testing.T) {
	fmt.Println(ParseLikeSql("sfsfds%_"))
}

func ParseLikeSql(s string) string {
	var escapes, keywords = `\`, []string{"%", "_"}
	for _, keyword := range keywords {
		if strings.Contains(s, keyword) {
			s = strings.Replace(s, keyword, escapes+keyword, -1)
		}
	}
	return s
}
