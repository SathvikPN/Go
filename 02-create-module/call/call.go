package main

import (
    "fmt"

    "example.com/greet" // mod file replace this with local directory
)

func main() {
    // Get a greeting message and print it.
    message1 := welcome.Hello("Gladys") // error
	msg2 := welcome.Hi("Sathvik")

    fmt.Println(message1)
	fmt.Println(msg2)
}