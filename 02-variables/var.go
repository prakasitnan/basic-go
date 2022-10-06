package main

import "fmt"

func main() {
	var name string
	s := "Gopher นำแน่"
	i := 5
	f := 3.7
	b := true
	r := 'ก'

	fmt.Printf("name: %q\n", name)
	fmt.Printf("name: %T\n", name)

	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("runr: %v\n", r)

}
