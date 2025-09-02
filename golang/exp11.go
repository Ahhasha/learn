package main

import "fmt"

func main() {
	var a, b, d int = 1, 2, 4

	p_nums := [4]*int{&a, &b, nil, &d}

	for _, p := range p_nums {
		if p != nil {
			fmt.Println(*p)
		}
	}
}
