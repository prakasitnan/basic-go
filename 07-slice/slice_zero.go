package main

import "fmt"

func main() {
	// langs := []string{"golang", "python", "javascript"}
	// langs := []string{}
	var langs []string
	fmt.Printf("langs : %#v\n", langs)

	if langs == nil {
		fmt.Println(`Yes nil "langs" is a nil slice`)
	} else {
		fmt.Printf("langs is NOT nil, value: %#v\n", langs)
	}
}
