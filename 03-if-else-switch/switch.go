package main

import (
	"fmt"
	"time"
)

/*
	golang not put break, golang auto break
	fallthrough continue pass case like continue
*/

func main() {
	today := time.Monday

	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("สวัสดีวัน %v อยู่แดงมีแฮงเดย์\n", today)
	}
}
