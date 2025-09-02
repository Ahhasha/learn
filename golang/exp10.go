package main

import "fmt"

func main() {

	var a int = 9

	p_a := &a
	p_p_a := &p_a

	fmt.Println("Value of a:", **p_p_a)
	fmt.Println("Value of p_a (Address of a):", *p_p_a)
	fmt.Println("Address of p_a:", p_p_a)

}
