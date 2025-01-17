package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
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
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"), // Replace with your desired region
				},
			}))
			ec2Svc := ec2.New(sess)
			input := &ec2.DescribeInstancesInput{
				Filters: []*ec2.Filter{
					{
						Name:   aws.String("instance-state-name"),
						Values: []*string{aws.String("running")},
					},
				},
			}

			result, err := ec2Svc.DescribeInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result)
			}

		case "2":
			fmt.Println("Listing S3 buckets...")
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"), // Replace with your desired region
				},
			}))
			s3Svc := s3.New(sess)
			result, err := s3Svc.ListBuckets(nil)
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Println("Success", result)
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
