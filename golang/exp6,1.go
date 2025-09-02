package main
import "fmt"

func main() {
	add(1, 2, 3.4, 5.6, 1.2)
}
func add(x, y int, a, b, c float32) {
	z := x + y
	d := a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}