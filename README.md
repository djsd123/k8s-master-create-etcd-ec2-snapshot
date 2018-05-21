# k8s-master-create-etcd-ec2-snapshot
Lambda function to create snapshots of etcd volumes attached to master nodes


The below might allow a nicer interface to the TagResource function

```
func TagResource(resourceID string, tags map[string]string) {
	tags := []*ec2.Tag{}
	for key, value := range tagsMap {
		tags = tags.append(
			ec2.Tag{
				Key:   aws.String(key),
				Value: aws.String(value),
			}
		)
	}


	tagsInput := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(resourceID),
		},
		Tags: tags,
	}

	// Finally tag snapshot with instance ID, Device designation and Volume name
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

```