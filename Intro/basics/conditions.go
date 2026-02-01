package main
import (
	"fmt"
)
func main() {
	   age := 18
	   if age < 18 {
	      fmt.Println("You are a minor.")
	   } else if age >= 18 && age < 65 {
	      fmt.Println("You are an adult.")
	   } else {
	      fmt.Println("You are a senior citizen.")
	   }
	   for i := 1; i <= 5; i++ {
	      fmt.Print("Iteration:!!", i," ")
		  // if I want to print in same line 

		}
		// while loop
		count :=1
		for count <= 5 {
			fmt.Println("Count:", count," ")
			count++
		}
	 // printf and print is different
	   fmt.Printf("\nThis is a formatted number: %.2f\n", 123.456789)
	   fmt.println("This is a simple print statement.")
	   fmt.print("This is another print statement without a newline.")	
}