package main

import (
	"fmt"
)

func main() {
	var fruits = []string{"Apple", "Grapes", "Orange"}
	fmt.Println("Fruits Array ", fruits)
	fruits = append(fruits, "Mango", "Banana")
	fmt.Println("Fruits modified", fruits)
	fmt.Println("Fruits include till 2 ", append(fruits[:2]))
	fmt.Println("Fruits exclude from 1 ", append(fruits[1:]))
	fruits = append(fruits[1:3]) // From To
	fmt.Println("Fruits slice", fruits)

	sports := make([]string, 3)
	fmt.Println("Sports = ", sports) // 3 empty strings
	sports[0] = "Hockey"
	sports[1] = "Football"
	fmt.Println("Sports defined indexes = ", sports) // 2 filled 1 empty
	sports[2] = "Cricket"
	// sports[3] = "Basketball"
	// fmt.Println("Sport at 3rd index added ", sports) panic: runtime error: index out of range [3] with length 3
	sports = append(sports, "Basketball", "Volleyball")
	fmt.Println("Sports ", sports) // This works (saves memory - performance optimisation)
	const index int = 2            // Removing element
	sports = append(sports[:index], sports[index+1:]...)
	fmt.Println(sports)
}
