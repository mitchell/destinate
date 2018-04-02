package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchelljfs/destinate/backend/schema"
)

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var re schema.Review

	err := re.Destroy(r.PathParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
