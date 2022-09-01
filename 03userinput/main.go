package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input Program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey! Choose on a scale of 1 to 10 where 10 be best and 1 be the worst: ")

	// comma ok || comma err
	// here underscore is acting as a catch if any error comes up
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving the valuable feedback, ", input)
}
