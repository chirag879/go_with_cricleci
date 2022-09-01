package main

import "fmt"

// Here in GoLang 1st letter capital signifies that it is a Public Variable
const LoginToken string = "kefohfuh383u939u9f"

func main() {
	var username string = "Chirag"
	fmt.Println(username)
	fmt.Printf("The data type of variable is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The data type of variable is: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("The data type of variable is: %T \n", smallInt)

	var smallFloat float32 = 251.11111111111111111
	fmt.Println(smallFloat)
	fmt.Printf("The data type of variable is: %T \n", smallFloat)

	// default values & some aliases
	var newVariable string
	fmt.Println(newVariable)
	fmt.Printf("The data type of variable is: %T \n", newVariable)

	// no var style -> := Walrus Operator(only use in a method global use is not allowed)
	withoutVar := 80000.01
	fmt.Println(withoutVar)

	fmt.Println(LoginToken)
	fmt.Printf("The data type of variable is: %T \n", LoginToken)
}
