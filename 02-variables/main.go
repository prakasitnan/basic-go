package main

import "fmt"

func main() {
	// var name string = "Gopher นำแน่!!!"
	// var name := "Gopher นำแน่!!!"
	// name := "Gopher นำแน่!!!"
	var name string
	s := "Gopher นำแน่!!!"
	i := 5
	f := 3.7
	b := true
	r := 'ก'

	fmt.Printf("name: %v\n", name)
	fmt.Printf("type: %T\n", name)

	fmt.Printf("int: %T\n", s)
	fmt.Printf("float64: %T\n", i)
	fmt.Printf("bool: %T\n", f)
	fmt.Printf("string: %T\n", b)
	fmt.Printf("rune: %T\n", r)
}
