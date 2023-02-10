package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi 🖐🖐, welcome to my quiz game!")

	fmt.Printf("Enter Your FullName: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Hello %v 🤍❤, welcome to the game! \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scanln(&age)

	if age >= 20 {
		fmt.Println("Yeay you can start the game ")
	} else {
		fmt.Println("Sorry you cannot play this game")
		return
	}

	scores := 0

	fmt.Printf("what is another word for Analyze? ")
	var answer1 string
	fmt.Scan(&answer1)

	if answer1 == "think" || answer1 == "THINK" {
		fmt.Println("You got the answer correct")
		scores += 10
	} else {
		fmt.Println("Sorry!, your answer is wrong")
	}

	fmt.Printf("What is the part of HT001 that stores the AC? ")
	var answer2 string
	fmt.Scan(&answer2)
	if answer2 == "COMPRESSOR" || answer2 == "compressor" {
		fmt.Println("You got the answer correct")
		scores += 1
	} else {
		fmt.Println("Sorry!, you got it wrong")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer3 uint
	fmt.Scan(&answer3)
	if answer3 == 4 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer31 uint
	fmt.Scan(&answer31)
	if answer31 == 46 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer4 uint
	fmt.Scan(&answer4)
	if answer4 == 48 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer5 uint
	fmt.Scan(&answer5)
	if answer5 == 42 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer6 uint
	fmt.Scan(&answer6)
	if answer6 == 40 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer7 uint
	fmt.Scan(&answer7)
	if answer7 == 41 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer8 uint
	fmt.Scan(&answer8)
	if answer8 == 47 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer9 uint
	fmt.Scan(&answer9)
	if answer9 == 44 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("How many parts has the AC of HT001? ")
	var answer10 uint
	fmt.Scan(&answer10)
	if answer10 == 55 {
		fmt.Println("Yay, you got it correct!")
		scores += 10
	} else {
		fmt.Println("sorry, you got it wrong!")
	}
	fmt.Printf("You scored: %v%% .\n", scores)
	if scores >= 90 {
		fmt.Println("Excellent, you passed the Quiz.An Email will be sent to you soon for more sessions.Thank You and Best Wishes")
	} else if scores >= 70 && scores <= 89 {
		fmt.Println("Congratulations, you passed the Quiz.An Email will be sent to you soon for more sessions.Thank You and Best Wishes")
	} else {
		fmt.Println("Sorry, you could not pass the quiz better luck next time. you can still check our newsletter for more information. Thank You")
	}

}
