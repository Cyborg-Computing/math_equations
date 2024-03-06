package main

import (
	"fmt"
	"math"
)

func quadraticFormula(a, b, c float64) (float64, float64) {
	x1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	x2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	return x1, x2
}

func main() {
	var a, b, c float64
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
	fmt.Print("Enter value for c: ")
	if _, err := fmt.Scan(&c); err != nil {
		fmt.Println("Error reading c:", err)
		return
	}
	x1, x2 := quadraticFormula(a, b, c)
	fmt.Printf("The roots are %f and %f\n", x1, x2)

	// Print if the roots are imaginary
	if math.IsNaN(x1) || math.IsNaN(x2) {
		fmt.Println("The roots are imaginary, you cannot square root a negative number")
	}
}
