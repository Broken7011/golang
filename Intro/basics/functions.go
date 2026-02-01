package main

import (
	"fmt" 
	"errors"
)
func main() {
	// hello()
    // add(23,56,100)
	quotient, remainder, err := calculate(7,0)
	if err != nil {
		fmt.Println("An error occured",err)
		return 
	}
	fmt.Printf("7 is divided by 3 is  %d with remainder and %d quotient\n",remainder,quotient)
}
// func hello(){
// 	fmt.Print("Hello")
// }
// func add(x int,y int , z int){
//    fmt.Println("Value of x",x)
//    fmt.Println("Value of y",y)
//    fmt.Println("Value of z",z)
//    ans := x+y+z
//    fmt.Println(ans)
// }
func calculate(dividend int, divisor int)(int,int,error){
	if divisor == 0{
		return 0,0, errors.New("Cannot divide zero")
	}
	quotient := dividend/divisor 
	remainder := dividend % divisor  
	return quotient,remainder,nil
 	
}