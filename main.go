package main

import (
	"fmt"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func showQuestions() {
	fmt.Println("Why do we use methods?")
	option1 := "1. To repeat a statement multiple times"
	option2 := "2. To decompose a program into several small subroutines"
	option3 := "3. To determine the execution time of a program"
	option4 := "4. To interrupt the execution of a program"

	var optionSlice []string
	optionSlice = append(optionSlice, option1, option2, option3, option4)

	for _, element := range optionSlice {
		fmt.Println(element)
	}
}

func getUserChoice() string {
	var choice string
	fmt.Scan(&choice)
	return choice
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("")
	// write your code here
	showQuestions()

	var userChoice = getUserChoice()
	for userChoice != "4" {
		fmt.Println("Please try again")
		showQuestions()
		userChoice = getUserChoice()
	}
	fmt.Println("Great! you did it 🚀")
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Gobot", "2022") // change it as you need
	showName()
	guessAge()
	count()
	// call something here
	startQuiz()
	sayGoodbye()
}
