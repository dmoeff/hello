package main

import (
	"fmt"

	"github.com/dmoeff/greet"
)

func main() {
	// Get a greeting message and print it.
	message := greet.Hello("Moeff")
	fmt.Println(message)
}
