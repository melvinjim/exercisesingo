package main

import (
	"calculator/models"
	"fmt"
)

func main() {
	n := models.New(29,37)
	fmt.Printf("----------------------\n")
	fmt.Printf("The result is: %d\n", (n.Multiplication()))
	fmt.Printf("----------------------\n")
	fmt.Printf("The result is: %d\n", (n.Division()))
	fmt.Printf("----------------------\n")
	fmt.Printf("The result is: %d\n", (n.Addition()))
	fmt.Printf("----------------------\n")
	fmt.Printf("The result is: %d\n", (n.Subtraction()))
}