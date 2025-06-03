package main
import "fmt"

// This is a placeholder for the main function
func main() {
	//fmt.Print("Sat, number 2, no funfact")
	// var student1 string
	// student1 = "Sat"
	// fmt.Print(student1)
	// var attended bool
	// attended = true
	// var name string
	// var luckyNumber int
	// var funFact string

	// name = "Sat"
	// luckyNumber = 2
	// funFact = "I love to play chess."
	// fmt.Println(name)
	// fmt.Println(luckyNumber)
	// fmt.Println(funFact)

	// var age int
	// fmt.Println("Enter your age:")
	// fmt.Scanln(&age)
	// fmt.Println("Your age is:", age)
	var name string
	var luckyNumber int
	var funFact string

	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("What is your lucky number?")
	fmt.Scanln(&luckyNumber)
	fmt.Println("Share something interesting about yourself:")
	fmt.Scanln(&funFact)
	fmt.Println("Your name is:", name)
	fmt.Println("Your lucky number is:", luckyNumber)
	fmt.Println("something interesting:", funFact)
}
