package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a resource to list:")
		fmt.Println("1. EC2")
		fmt.Println("2. S3")
		fmt.Println("3. Lambda")
		fmt.Println("4. Exit")
		fmt.Print("Enter the number of your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim and process input
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("Listing EC2 instances...")
		case "2":
			fmt.Println("Listing S3 buckets...")
			// Add logic to list S3 buckets here
		case "3":
			fmt.Println("Listing Lambda functions...")
			// Add logic to list Lambda functions here
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
