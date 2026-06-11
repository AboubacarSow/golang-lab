package main

import "os"
func callbackExit(c *config) error{
	os.Exit(0)
	return nil
}