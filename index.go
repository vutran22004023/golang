package main

import (
	"fmt"
	"learing/learinggo"
)

func main() {
	price, no := 90, 6
	totalPrice := learinggo.CalculateBill(price, no) // corrected the package and function name
	fmt.Println("Hello, World!", totalPrice)
}
