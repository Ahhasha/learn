package main
import "fmt"

func main() {
	add(4, 5)
	add(20, 6)
}

func add(x int, y int){
	z := x + y
	fmt.Println("x + y = ", z)
}