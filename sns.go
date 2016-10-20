package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"os"
)

// SNS class for handling Simple Notification Work
type SNS struct{}

// main - use with the following flags
/*
 -s the flag to set up a topic. Pass the topic name and aws region, in that order.
 -r the flag to register a device. Pass the device token, platformARN and aws region, in that order.
 -p the flag to push a notification. Pass the deviceARN (target), text and aws region, in that order.
*/
func main() {

	if os.Args[1] == "-s" {
		if len(os.Args) < 4 {
			println("not enough args")
			return
		} else {
			setUpSNS(os.Args[2], os.Args[3])
		}
	} else if os.Args[1] == "-r" {
		if len(os.Args) < 5 {
			println("not enough args")
			return
		} else {
			registerDevice(os.Args[2], os.Args[3], os.Args[4])
		}
	} else if os.Args[1] == "-p" {
		if len(os.Args) < 5 {
			println("not enough args")
			return
		} else {
			pushNotification(os.Args[2], os.Args[3], os.Args[4])
		}
	} else {
		fmt.Println(`please pass a valid flag :\n
		 -s to set up a topic. Pass the topic name, aws region, in that order.\n
		 -r to register a device. Pass the device token, platformARN, aws region, in that order.\n
		 -p to push a notification. Pass the deviceARN, text, aws region, in that order.`)
	}
}

// setUpSNS sets up the topic and preps everything for the push system for use subscriptions;
func setUpSNS(name string, region string) error {

	svc := sns.New(session.New(), &aws.Config{Region: aws.String(region)})

	params := &sns.CreateTopicInput{
		Name: aws.String(name), // Required
	}
	resp, err := svc.CreateTopic(params)

	log.Println(resp)
	log.Println(err.Error())

	return err
}

// registerDevice adds a device token to the ARN
func registerDevice(token string, platformARN string, region string) (*string, error) {

	svc := sns.New(session.New(), &aws.Config{Region: aws.String(region)})

	params := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(platformARN),
		Token:      aws.String(token),
		Attributes: nil,
	}

	resp, err := svc.CreatePlatformEndpoint(params)

	log.Println(resp)
	log.Println(err.Error())

	return resp.EndpointArn, err
}

// pushNotification sends the notification
func pushNotification(arn string, text string, region string) error {

	svc := sns.New(session.New(), &aws.Config{Region: aws.String(region)})

	input := &sns.PublishInput{
		Message: aws.String(text), // Required
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"Key": { // Required
				DataType:    aws.String("String"), // Required
				StringValue: aws.String(text),
			},
		},
		TargetArn: aws.String(arn),
	}

	result, err := svc.Publish(input)

	log.Println(result)
	log.Println(err.Error())

	return err
}
