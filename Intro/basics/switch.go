package main

import "fmt"

func main() {
	var guess int
	for {
		fmt.Print("Enter guess no: ")
		fmt.Scan(&guess)
		if guess >99{
			fmt.Println("Guess is too high")
		}else if guess < 99{
			fmt.Println("Guess is too low")
		}else{
			fmt.Println("Congrats!! you guess it right!!")
			break
		}
	}

}