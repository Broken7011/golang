// user-defined type that groups reated data structs
package main
import "fmt"

// defining a struct type

func main() {
	// creating an instance of struct
	// var p1 Person	
	// p1.name = "Rishabh"
	// p1.age = 25
	// p1.country = "India"	
	type Person struct {
	name   string	
	age    int
	country string
}
	rishabh := Person{name: "Rishabh", age: 25, country: "India"}
	fmt.Println("Person 1:", rishabh.age)

   // we can also create struct without defining type
	employee := struct {
		id    int
		name  string
		salary float64
	}{
		id:    101,
		name:  "Alice",
		salary: 75000.50,
	}
	fmt.Println("Employee Name:", employee.name)

	// But the best practice is to define struct type first and then create instances of it
	type Car struct {
		brand string
		model string
		year  int
	}
	myCar := Car{brand: "Toyota", model: "Camry", year: 2020}
	fmt.Println("My Car Model:", myCar.model)
}