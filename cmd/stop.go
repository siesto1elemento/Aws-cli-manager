package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var stopcmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a resource to Stop:")
		fmt.Println("1. EC2")
		fmt.Println("2. S3")
		fmt.Println("3. Lambda")
		fmt.Println("4. Exit")
		fmt.Print("Enter the number of your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)

		}
		choice := strings.TrimSpace(input)
		switch choice {
		case "1":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			}))
			ec2Svc := ec2.New(sess)
			instanceIDs := []*string{
				aws.String("ec2_instance_id"),
			}

			input := &ec2.TerminateInstancesInput{
				InstanceIds: instanceIDs,
			}

			result, err := ec2Svc.TerminateInstances(input)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("Result", result)
		case "2":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			}))
			s3Svc := s3.New(sess)

			input := &s3.DeleteBucketInput{
				Bucket: aws.String("Your_bucket_name"),
			}

			result, err := s3Svc.DeleteBucket(input)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("Result", result)
		case "3":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			}))
			lambdaSvc := lambda.New(sess)

			input := &lambda.DeleteFunctionInput{
				FunctionName: aws.String("Your_function_name"),
			}

			result, err := lambdaSvc.DeleteFunction(input)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("Result", result)

		}

	},
}

func init() {
	rootCmd.AddCommand(stopcmd)
}
