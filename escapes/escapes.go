package main

import "fmt"

// go run -gcflags '-m -l' escapes.go
type P struct {
	Age  int
	name *string
}
type UserData struct {
	Name string
	P    P
}

var i int

func NewData() *UserData {
	return &UserData{}
}

func main() {
	info := NewData()
	info.Name = "WilburXu"

	info = NewData()
	GetUserInfo(*info)
	Getp(info.P)

}

func GetUserInfo(userInfo UserData) {
	userInfo.P = P{Age: 99}
	var name = "www"
	userInfo.P.name = &name

}

func Getp(p P) {
	fmt.Println(i)
	fmt.Println(p.Age)
}
