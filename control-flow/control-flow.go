package main

import (
	"fmt"
)

func main() {

	// if true {
	// 	fmt.Println("Test passed")
	// }
	
	
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 		"California": 39123456,
	// 		"Texas": 4652123,
	// 		"Florida": 4564646,
	// 		"New York": 789456,
	// 		"Pennsylvania": 3456456,
	// 		"Illinois": 7954123,
	// 		"Ohio": 45123,
	// }
	// // Initializer Syntax
	// if pop, ok := statePopulations["Florida"]; ok {
	// 	//pop only defined in this scope
	// 	fmt.Println(pop)
	// }

	number := 50
	guess := 30
	//short circuiting. The runtime will not execute returnTrue
	// if guess <=1 || returnTrue() || guess > 100 {
	if guess <=1 {
		fmt.Println("the guess must be greater than 1!")

	} else if guess > 100 {
		fmt.Println("the guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
	
		if guess > number {
			fmt.Println("too hight")
		}
	
		if guess > number {
			fmt.Println("you got it!")
		}
	}

	switch 2 {
	case 1, 5, 10:
			fmt.Println("One")
			// switches need a break, but can have it
	case 2: 
			fmt.Println("two")
	default:
			fmt.Println("not one or two")
	}

	i := 10

	switch {
	case i <= 10:
			fmt.Println("Less than or equal to ten")
	case i <= 20: 
			fmt.Println("less than or equal to twenty")
	default:
			fmt.Println("greater than twenty")
	}

	var j interface{} = 1

	// type switch. You can use types as the cases.
	// very useful if you need to do something diferent based on the parameter of a function
	switch j.(type) {
	case int:
			fmt.Println("i is an int")
	case float64:
			fmt.Println("i is a float64")
	case string:
			fmt.Println("i is a string")
	default:
			fmt.Println("i is another type")
	}

}

func returnTrue() bool {
	fmt.Println("Returing true")
	return true
}