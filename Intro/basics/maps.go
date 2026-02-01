package main
import "fmt"

// maps are used to store key-value pairs
// similar to dictionaries in python
// var m map[string]int  // declaring a map
// m = make(map[string]int) // initializing a map
// DOes not maintain order of elements AND keys are unique
func main() {
	contacts := map[string]string{
		"Rishabh": "1234567890",
		"Alice":   "9876543210",
		"Bob":     "5555555555",
	}
    //  empty map
	scores := make(map[string]int)  // make function to create a map 
	scores["Math"] = 95
	scores["Science"] = 90
	scores["English"] = 85
	fmt.Println("Scores:", scores)


	fmt.Println("Contacts:", contacts)
	contacts["Charlie"] = "1112223333" // adding new key-value pair
	fmt.Println("Updated Contacts:", contacts)
	
	delete(contacts, "Alice") // deleting a key-value pair
	fmt.Println("After deletion:", contacts)
	
	number := contacts["Rishabh"] // accessing value by key
	fmt.Println("Rishabh's number:", number)
	
	// checking if key exists
	value, exists := contacts["Bob"]
	if exists {
		fmt.Println("Bob's number:", value)
	} else {
		fmt.Println("Bob's number not found")
	}
}