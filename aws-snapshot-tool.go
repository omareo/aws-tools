package main

import (
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-west-2")}) //use the correct AWS Zone

	volumeParams := &ec2.DescribeVolumesInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag-value"),
				Values: []*string{
					aws.String("Backup"), //this is tagged in AWS
				},
			},
		},
	}

	volumeResponse, volumeError := svc.DescribeVolumes(volumeParams)
	if volumeError != nil {
		panic(volumeError)
	}


	for idx, res := range volumeResponse.Volumes {

	   timestamp := time.Now().Format(time.RFC850)

	   snapshotParams := &ec2.CreateSnapshotInput{
	       VolumeId:    aws.String(*res.VolumeId), 
	       Description: aws.String(timestamp),
	       DryRun:      aws.Bool(false),
	   }

	   snapshotResponse, snapshotError := svc.CreateSnapshot(snapshotParams)

	   if snapshotError != nil {
	       fmt.Println(snapshotError.Error())
	       return
	   }

	   fmt.Println(snapshotResponse)
	   fmt.Println(idx)
	}
}
