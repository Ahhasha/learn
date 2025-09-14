package main

import "fmt"

type Movable interface {
	moveX(distance int)
	moveY(distance int)
}

type Rectangle struct{}

func (r Rectangle) moveX(distance int) {
	fmt.Println("Перемещаем прямоугольник на", distance, "по оси X")
}

func (r Rectangle) moveY(distance int) {
	fmt.Println("Перемещаем прямоугольник на", distance, "по оси Y")
}

func move_object(obj Movable) {
	obj.moveX(10)
}

func main() {
	rect := Rectangle{}
	move_object(rect)
}
