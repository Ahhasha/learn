package main

import "fmt"

type Reader interface {
	read()
}
type File struct {
	text string
}

func (f File) read() {
	fmt.Println(f.text)
}
func read_data(data Reader) {
	data.read()
}

func main() {
	file := File{"Hello go runner"}
	read_data(file)

	p_file := &file
	read_data(p_file)
}
