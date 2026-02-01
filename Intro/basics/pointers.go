// pointers in go variable that stores the memory address of another vairable 
// used to directly access and manipulate memory locations
package main
import "fmt"	

func main() {
	var a int = 42
	var p *int = &a // p is a pointer to an integer, &a gives the address of a	
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value at address stored in p:", *p) // dereferencing pointer to get value
	*p = 100 // updating value at address stored in p
	fmt.Println("Updated value of a:", a)
}