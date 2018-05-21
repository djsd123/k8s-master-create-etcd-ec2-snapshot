package tags

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var connection = ec2.New(session.New())

func FetchResourceTags(resourceID *string) (getTagsOutput *ec2.DescribeTagsOutput) {

	volTagsInput := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []*string{
					aws.String(*resourceID),
				},
			},
		},
	}
	// Get the volume's tags
	getTagsOutput, err := connection.DescribeTags(volTagsInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

	return getTagsOutput

}

func TagResource(resourceID *string, tags []*ec2.Tag) {

	tagsInput := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(*resourceID),
		},
		Tags: tags,
	}

	tagIt, err := connection.CreateTags(tagsInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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
	fmt.Println(tagIt)

}
