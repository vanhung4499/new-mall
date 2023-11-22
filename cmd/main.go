package cmd

import "fmt"

func main() {
	loading()
	fmt.Println("Start successfully...")
}

func loading() {
	fmt.Println("Loading configuration completed...")
	go scriptStarting()
}

func scriptStarting() {
	// start some scripts
}
