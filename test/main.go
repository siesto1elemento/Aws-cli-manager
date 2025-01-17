package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Create a new session with the shared config enabled
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	// Retrieve credentials from the default provider chain
	creds := sess.Config.Credentials

	// Attempt to retrieve the credentials
	value, err := creds.Get()
	if err != nil {
		log.Fatalf("Unable to retrieve AWS credentials: %v", err)
	}

	// Print the credentials to verify
	fmt.Println("AWS Access Key ID:", value.AccessKeyID)
	fmt.Println("AWS Secret Access Key:", value.SecretAccessKey)
}
