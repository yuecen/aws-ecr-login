package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	// Create a ECR client with additional configuration
	region := flag.String("region", "<empty>", "for example: us-east-1")
	flag.Parse()
	svc := ecr.New(session.New(), aws.NewConfig().WithRegion(*region))

	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	authData := result.AuthorizationData[0]
	endpoint := string(*authData.ProxyEndpoint)
	// A base64-encoded string that contains authorization data for the specified
	// Amazon ECR registry. When the string is decoded, it is presented in the
	// format user:password for private registry authentication using docker
	// login.
	// https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_AuthorizationData.html
	token, err := base64.StdEncoding.DecodeString(*authData.AuthorizationToken)
	re := regexp.MustCompile(`(?P<user>.+):(?P<password>.+)`)
	parsed := re.FindStringSubmatch(string(token))
	user := parsed[1]
	password := parsed[2]

	fmt.Printf("docker login -u %s -p %s %s\n", user, password, endpoint)
}
