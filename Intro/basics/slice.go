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
	array4 = append(array4)
	// SLices are same as of arrays but dynamic in size
	// we can use append function to add more elements to slice
	// Slices do not have fixed size
	num := make([]int,4,5) // length 4, capacity 5
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num))
	// make function is used to create slices, maps and channels 
	// first argument is type, second is length, third is capacity
	// length is number of elements in slice
	// capacity is total space allocated for slice
	// for example if we create slice of length 4 and capacity 5
	num = append(num,10)
	fmt.Println("After appending:")
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num)) // capacity will remain same until length exceeds capacity
	num = append(num,20)
	fmt.Println("After appending again:")
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num)) // capacity will double when length exceeds capacity
	// capacity is automatically managed by Go runtime as we append elements to slice 

	// creating slice from array
	arr := [6]string{"A","B","C","D","E","F"}
	slice1 := arr[1:4] // slice from index 1 to 3
	fmt.Println("Slice1:", slice1)
	slice2 := arr[:3] // slice from start to index 2
	fmt.Println("Slice2:", slice2)
	slice3 := arr[2:] // slice from index 2 to end
	fmt.Println("Slice3:", slice3)
}