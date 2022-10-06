package main

import "fmt"

func main() {
	fmt.Println("hello Gopher!!!")

	fmt.Printf("hello %s!!!\n", "Gopher")
	fmt.Println("next line")

	fmt.Printf("hello %d!!!\n", 66)

	fmt.Printf("hello %v!!!\n", "Gopher")
	fmt.Printf("hello %v!!!\n", 66)

	fmt.Printf("hello %v : %v!!!\n", "Gopher", 55)
}
