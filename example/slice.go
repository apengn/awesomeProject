package main

import "fmt"

// 如果当前slice大小小于1024，按每次2倍增长，否则每次按当前大小1/4增长。直到增长的大小超过或等于新大小。
// func main() {
//
// 	slice := make([]int, 0, 0)
//
// 	for i := 0; i < 22; i++ {
// 		slice = append(slice, i)
// 		fmt.Println("========================")
// 		fmt.Printf("len:%d\n", len(slice))
// 		fmt.Printf("cap:%d\n", cap(slice))
// 		fmt.Println(slice)
//
// 	}
//
// }

//
func main() {

	type S struct {
		Test []string
	}

	s := S{}

	s.Test = append(s.Test, "4")
	s.Test = append(s.Test, "2")
	s.Test = append(s.Test, "5")

	fmt.Printf("len:%d\n", len(s.Test))
	fmt.Printf("cap:%d\n", cap(s.Test))
	s2 := make([]string, len(s.Test), cap(s.Test)*2)
	copy(s2, s.Test)
	fmt.Printf("len:%d\n", len(s2))
	fmt.Printf("cap:%d\n", cap(s2))
	var s3 []string
	s3 = nil

	fmt.Println(s3)

	fmt.Printf("len:%d\n", len(s3))
	fmt.Printf("cap:%d\n", cap(s3))
	// fmt.Println(s.Test)
}
