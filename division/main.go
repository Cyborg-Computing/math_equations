package main

import (
	"fmt"
)

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var a, b float64
	fmt.Print("Enter value for a: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Error reading a:", err)
		return
	}
	fmt.Print("Enter value for b: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Error reading b:", err)
		return
	}
	result, err := division(a, b)
	if err != nil {
		fmt.Println("Error dividing:", err)
		return
	}
	fmt.Printf("The answer is %f\n", result)
}
