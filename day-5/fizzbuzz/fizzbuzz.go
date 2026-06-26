package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter the number-: ")
	fmt.Scan(&input)
	if input%5 == 0 && input%3 == 0 {
		fmt.Print("FizzBuzz\n")
	} else if input%3 == 0 {
		fmt.Print("Fizz\n")
	} else if input%5 == 0 {
		fmt.Print("Buzz\n")
	} else {
		fmt.Printf("%v is not divisible by neither 3 nor 5", input)
	}
}
