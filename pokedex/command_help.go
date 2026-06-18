package main

import "fmt"

func callbackHelp(c *config,args ...string) error{
	fmt.Println("Welcome to Pokedex help menu")
	fmt.Println("Here are available commands: ")
	commands:=getCommands()
	for _, cmd:= range commands{
		if cmd.name=="invalid" {
			continue
		}
		fmt.Printf("- %s: %s\n",cmd.name, cmd.description)
	}

	return nil
	
}