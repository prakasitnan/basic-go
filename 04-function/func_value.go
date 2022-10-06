package main

import "fmt"

func say(n string) {
	fmt.Printf("My name is %s\n", n)
}

func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println("result :\n", r)
}

func main() {
	var name = "Nong"
	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)
	say(name)

	var plus = func(a int, b int) int { return a + b }

	p := plus(1, 2)
	fmt.Println("1+2 =", p)
	fmt.Printf("type: %T\n", plus)

	cal(plus)

	minus := func(a, b int) int { return a - b }
	cal(minus)
}
