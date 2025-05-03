package main

import "log"

func main() {
	// this will block pushing to the repo by the pre push hook
	return

	log.Println("Hello, World!")
	log.Println("This is a simple Go program for Github actions.")
}
