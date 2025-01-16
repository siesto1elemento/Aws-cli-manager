package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-cli-manager",
	Short: "A CLI to manage AWS resources",
	Long:  `A command-line interface tool to interact with AWS services such as EC2, S3, IAM, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to aws-cli-manager!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
