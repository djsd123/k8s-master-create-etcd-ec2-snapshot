package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
	"./tags"
	"./snapshot"
)

var connection = ec2.New(session.New())


func main() {

	// Filter instances with role set to master
	instanceKeyFilter   := os.Getenv("INSTANCE_TAG_KEY")
	instanceValueFilter := os.Getenv("INSTANCE_TAG_VALUE")

	instanceInput := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String(fmt.Sprintf("tag:%s", instanceKeyFilter)),
				Values: []*string{aws.String(instanceValueFilter)},
			},
		},
	}

	// A reservation corresponds to a command to start instances
	// A reservation is what you do to provision instances, while an instance is what you get
	results, err := connection.DescribeInstances(instanceInput)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, res := range results.Reservations {

		for _, instance := range res.Instances {

			for _, blk := range instance.BlockDeviceMappings {

				// Exclude root volumes.  Root volumes have "Delete on termination" set to true mostly
				isRootDevice := *blk.Ebs.DeleteOnTermination

				if isRootDevice != true {
					vol_id := blk.Ebs.VolumeId

					// Perform and tag the snapshot
					snapShot := create_volume_snapshot.CreateSnapshot(vol_id)

					getVolumesTags := tags.FetchResourceTags(vol_id)

					snapShotID   := *snapShot.SnapshotId
					instanceName := *instance.Tags[1].Value
					deviceName   := *blk.DeviceName
					volumeName   := *getVolumesTags.Tags[1].Value

					Tags := []*ec2.Tag{
						{
							Key:   aws.String("Instance_Name"),
							Value: aws.String(instanceName),
						},
						{
							Key:   aws.String("Device_Name"),
							Value: aws.String(deviceName),
						},
						{
							Key:   aws.String("Name"),
							Value: aws.String(volumeName),
						},
					}
					tags.TagResource(snapShotID, Tags)
				}
			}
		}
	}
}
