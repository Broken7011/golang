package main

import "fmt"

func main() {
   	result := sum(10, 20, 30, 40, 50)
	fmt.Println("Sum:", result) // Output: Sum: 150	
}
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}