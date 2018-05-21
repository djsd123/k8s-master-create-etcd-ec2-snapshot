package create_volume_snapshot

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
)

var connection = ec2.New(session.New())

func CreateSnapshot(volumeID *string) (snapshot *ec2.Snapshot) {

	snapshotInput := &ec2.CreateSnapshotInput{
		VolumeId: volumeID,
	}

	// Perform ebs snapshot operation
	snapIt, err := connection.CreateSnapshot(snapshotInput)
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
	fmt.Println(snapIt)

	return snapshot

}
