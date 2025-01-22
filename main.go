package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/siesto1elemento/Aws-cli-manager/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like to do?")
	fmt.Println("1. List AWS resources")
	fmt.Println("2. Stop AWS resources")
	fmt.Println("3. Delete AWS resources")
	fmt.Println("4. Create AWS resources")
	fmt.Println("5. Exit")
	fmt.Print("Enter the number of your choice: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		// Call the 'list' command
		os.Args = []string{"aws-cli-manager", "list"} // Simulate "aws-cli-manager list"
		cmd.Execute()
	case "2":
		// Call the 'delete' command
		os.Args = []string{"aws-cli-manager", "stop"} // Simulate "aws-cli-manager delete"
		cmd.Execute()
	case "3":
		// Call the 'delete' command
		os.Args = []string{"aws-cli-manager", "delete"} // Simulate "aws-cli-manager delete"
		cmd.Execute()
	case "4":
		os.Args = []string{"aws-cli-manager", "create"}
		cmd.Execute()
	case "5":
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
