package main
import (
	"fmt"
)

func main() {
   fmt.Println("Hello My name is Lala")	
   name := "Rishabh"
   age := 21
   salary := 50000.50
   isEmployed := true
   fmt.Println("Name:", name) 
   fmt.Println("Name: ", name)
   fmt.Println("Age:", age)
   fmt.Println("Salary:", salary)
   fmt.Println("Employed:", isEmployed)
   fmt.Println("My name is", name, "and my age is", age)
   fmt.Println(name+" is "+string(age)+" years old.")
   const(
      country = "India"
	  city = "New Delhi"
   )
   var (
	  Population = 1393409038
   )
   fmt.Println("Country:", country)
   fmt.Println("City:", city)
   fmt.Println("Population:", Population)
}