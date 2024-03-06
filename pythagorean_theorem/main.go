package main

import (
	"fmt"
	"math"
)

func pythagoreanTheorem(a, b float64) float64 {
	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return hypotenuse
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

	hypotenuse := pythagoreanTheorem(a, b)
	fmt.Printf("The hypotenuse is %f\n", hypotenuse)
}
