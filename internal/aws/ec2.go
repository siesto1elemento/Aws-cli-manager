package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/siesto1elemento/Aws-cli-manager/internal/state"
)

func ProvisionEC2(resource *state.Resource) error {
	ami := resource.Config["ami"].(string)
	instanceType := resource.Config["instance_type"].(string)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	ec2Svc := ec2.New(sess)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(ami),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	result, err := ec2Svc.RunInstances(input)
	fmt.Println("result", result)
	return err
}
