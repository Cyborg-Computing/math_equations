package main

import (
	"fmt"
	"math"
)

// create a program that calculates the roots of a quadratic equation
func quadraticFormula(a, b, c float64) (float64, float64, error) {
	if a == 0 {
		return 0, 0, fmt.Errorf("coefficient a cannot be zero")
	}
	discriminant := math.Pow(b, 2) - (4 * a * c)
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("negative discriminant, the roots are imaginary, you cannot square root a negative number")
	}
	x1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	x2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)

	return x1, x2, nil
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
	x1, x2, err := quadraticFormula(a, b, c)
	if err != nil {
		fmt.Println("Error calculating the roots:", err)
		return
	}
	fmt.Printf("The roots are %f and %f\n", x1, x2)

}
