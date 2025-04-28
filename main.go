package main

import (
	"dog-tracking/data"
	"fmt"
)

func main() {
	fmt.Println("What type of API do you want to use?\n - 1 - net/http\n - 2 - gin")
	var s string
	fmt.Scanln(&s)

	switch s {
	case "1":
		data.SetupHttp("8080")
	case "2":
		data.SetupGin("8080")
	default:
		fmt.Println("Choice not available")
	}
}
