// ARRAY is fixed size, SLICE is dynamic size
// size must be known at compile time for ARRAY
// once created, size of ARRAY cannot be changed
// var a [5]int  // array of 5 integers 
package main
import "fmt"
func main() {
	// var a [5]int  // will print 0 as default value
	// scores := [5]int{90, 80, 70, 60, 50}
	// fmt.Println("Array a:", a)
	// fmt.Println("Array scores:", scores) 
	// first_score := scores[0]
	// last_score := scores[4]
	// fmt.Println("First score:", first_score)
	// fmt.Println("Last score:", last_score)
	// scores[2] = 75 // updating value at index 2
	// fmt.Print(scores[4])

	friends := [5]string{"Rishabh1","Rishabh2","Rishabh3","Rishabh4","Rishabh5"}
	for i := 0; i < 5; i++ {
		fmt.Println(friends[i])
	}
	array2 := [...] int{1,2,3,4,5}
	array3 := [] int{1,2} // slice 
 	fmt.Println(len(array3))
	fmt.Println(len(array2))
	array4 := [] string{"Apple","Mango","Pineapple"}
	fmt.Println(len(array4))
	array3 = append(array3,6)
	fmt.Println(len(array3))
	// All 
	array4 = append()
	
}