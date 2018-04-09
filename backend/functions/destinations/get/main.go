package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// response, err := d.Find(r.PathParameters["id"])
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		Body:       err.Error(),
	// 		StatusCode: 400,
	// 	}, nil
	// }

	return events.APIGatewayProxyResponse{
		Body:       "Under Construction.",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
