package main

import "fmt"

func callbackInvalid(c *config) error {
	fmt.Println("Invalid command!")
	fmt.Println("---- Here are available commands: ------")
	commands := getCommands()
	for _, cmd := range commands {
		if cmd.name == "invalid" {
			continue
		}
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	return nil

}