package main

import (
	"fmt"
)

type Example struct {
	E *int32
}

func (in *Example) DeepCopyInto(out *Example) {
	*out = *in
	fmt.Println("*========")
	fmt.Println(*out)
	if in.E != nil {
		in, out := &in.E, &out.E
		if *in == nil {
			*out = nil
		} else {
			*in = new(int32)
			fmt.Println("*=======")
			fmt.Println(*in)
			**out = **in
			fmt.Println("**=======")
			fmt.Println(**out)
		}
	}

}
func main() {

	example := &Example{}
	var i int32 = 44
	example.E = &i

	out := &Example{}
	example.DeepCopyInto(out)

	fmt.Println(out.E)

}
