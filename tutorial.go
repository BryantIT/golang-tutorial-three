package main

import "fmt"

func main() {
	number := 260
	number = number + 5

	trueValue := true

	fmt.Printf("%T ", number)

	fmt.Println("Hello World!", number)
	fmt.Println("Truth ", trueValue)
}
