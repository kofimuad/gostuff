package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz game!")

	fmt.Println("Enter your name pls:")
	var name string
	fmt.Scan(&name)

	fmt.Println("Enter your age pls:")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("You're welcome %v, you are %v years old and eligible for the game\n", name, age)
	} else {
		fmt.Printf("You're welcome %v, you are %v years old hence you're not eligible for the game\n", name, age)
		return
	}

	fmt.Printf("What is the name of the first president of the republic of Ghana: ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	switch answer + " " + answer2 {
	case "Kwame Nkrumah":
		fmt.Printf("You're correct!")
	case "kwame nkrumah":
		fmt.Printf("You're correct!")
	default:
		fmt.Printf("Try again, you're wrong")
	}
}
