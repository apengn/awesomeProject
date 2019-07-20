package main

import "fmt"

type A struct {
	Name string
	B
}

type B struct {
	Age string
}
type C struct {
	A
	Address string
}

type D struct {
	C
	Year string
}

func (b *B) T() {
	fmt.Println(b.Age)
}

func (a *C) SetName(name string) {
	a.T()
	a.Name = name
}

func main1() {

	// 如果是0开始的数字为8进制，在进行计算的时 会自动转换为10进制 后在计算
	const (
		Century = 100
		Decade  = 010 // 8进制
		Year    = 0012
	)
	// The world's oldest person, Emma Morano, lived for a century,
	// two decades and two years.

	fmt.Println(2 * Year)
	fmt.Println(2 * Decade)
	fmt.Println("She was", Century+2*Decade+2*Year, "years old.")
}

func main2() {
	var b byte
	for b = 250; b <= 255; b++ {
		fmt.Printf("%d %c\n", b, b)
	}
}

func main3() {
	fmt.Println(5 ^ 2)
	fmt.Println(3^2+4^2 == 5^2)
	fmt.Println(10 ^ 2)
	fmt.Println(6^2+8^2 == 10^2)
}

func main4() {
	// i:=0
	// ++i
}

func main6() {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src) // Copy elements to dst from src.
	fmt.Println("dst:", dst)
	// why  output:
	// dst:[]

}

// 模拟3目运算    true?a:b
func Eye3(b bool, whenTrue, whenFalse interface{}) interface{} {
	if b {
		return whenTrue
	}
	return whenFalse
}

func main() {
	s := []byte("hello")
	s[0] = 'H'
	fmt.Println(string(s))
}
