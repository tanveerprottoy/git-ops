package main

import (
	"log"

	"github.com/tanveerprottoy/git-ops/mathext"
)

func main() {
	log.Println("Hello, World!")
	log.Println("This is an exercise project for GitOps")

	// add two numbers using the mathext package
	log.Println("Adding 7 + 4 =", mathext.Add(7, 4))
}
