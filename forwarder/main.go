package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.CloudwatchLogsEvent) error {
	cloudwatchLogsData, err := request.AWSLogs.Parse()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, e := range cloudwatchLogsData.LogEvents {
		fmt.Printf("Logged data is: %v", e)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
