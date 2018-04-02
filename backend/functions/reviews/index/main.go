package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchelljfs/destinate/backend/schema"
)

func handler() (events.APIGatewayProxyResponse, error) {
	var re schema.Review

	response := re.All()

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
