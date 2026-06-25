package main

import "fmt"

func main() {
	fmt.Println("Counting ")
	defer fmt.Println("done.")
	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}
}
