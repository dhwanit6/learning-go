package main

import (
	"fmt"
)

func orderfood() (string, string) {
	return "cheesburger", "success"
}

func main() {
	meal, status := orderfood()

	fmt.Println("I got a : ", meal)
	fmt.Println("The status is : ", status)
}
