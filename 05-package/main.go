package main

import (
	"fmt"

	"github.com/prakasitnan/igapp/time"
	"github.com/prakasitnan/igapp/user"
)

func main() {
	user.Info("Nong", "โกเฟอร์", 15)

	t := time.Today()
	fmt.Println("today is ", t)
}
