package main

import "fmt"

// This is a placeholder for the main function
func main() {
	var examResults int
	var ifPassed bool

	fmt.Println("what was your exam result?")
	fmt.Scanln(&examResults)

	if examResults > 50 {
		ifPassed = true
		fmt.Print("you passed the exam", ifPassed)
	} else {
		ifPassed = false
		fmt.Print("you did not pass the exam", ifPassed)
	}
	fmt.Println("passed:", ifPassed)
}
