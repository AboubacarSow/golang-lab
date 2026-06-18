package main

import "os"
func callbackExit(c *config,args ...string) error{
	os.Exit(0)
	return nil
}