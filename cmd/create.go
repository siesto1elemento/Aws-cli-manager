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

var createcmd = &cobra.Command{
	Use:   "create",
	Short: "create AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a resource to create:")
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

		choice := strings.TrimSpace(input)
		switch choice {
		case "1":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			}))
			ec2Svc := ec2.New(sess)
			input := &ec2.RunInstancesInput{
				ImageId:      aws.String("ami-a0cfeed8"),
				InstanceType: aws.String("t2.micro"),
				MinCount:     aws.Int64(1),
				MaxCount:     aws.Int64(1),
			}
			result, err := ec2Svc.RunInstances(input)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("Result:", result)
		case "2":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				}}))
			s3Svc := s3.New(sess)
			input := &s3.CreateBucketInput{
				Bucket: aws.String("rohitnewbucket9720"),
			}

			result, err := s3Svc.CreateBucket(input)
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Println("Result:", result)
		case "3":
			sess := session.Must(session.NewSessionWithOptions(session.Options{
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				}}))
			lambdaSvc := lambda.New(sess)

			input := &lambda.CreateFunctionInput{
				FunctionName: aws.String("my-dummy-function"),
				Runtime:      aws.String("nodejs18.x"),
				Handler:      aws.String("index.handler"),
				Role:         aws.String("your_arn_string"),
				Code: &lambda.FunctionCode{
					ZipFile: []byte("dummy"),
				},
			}

			result, err := lambdaSvc.CreateFunction(input)
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Println("Result", result)
		}

	},
}

func init() {
	rootCmd.AddCommand(createcmd)
}
