package main

import (
	"fmt"        // Formatting library
	"mockc/repl" // Our REPL
	"os"         // Operating system library
	"os/user"    // User package from the OS library
)

func main() {
	user, err := user.Current() // Returns current user

	if err != nil { // If any error is present, panic immediately!
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Moxie programming language!\n",
		user.Username)
	fmt.Printf("To start using it, just start typing in commands\n")
	repl.Start(os.Stdin, os.Stdout) // Assuming this is equiv. to Java System.stdin and System.stdout
}




