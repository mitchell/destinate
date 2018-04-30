package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchelljfs/destinate/backend/schema"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var re events.APIGatewayProxyResponse

	resp, err := schema.FindDestination(ctx, r.PathParameters["id"])
	if err != nil {
		re.Body = err.Error()
		re.StatusCode = 400
		return re, nil
	}

	jsonba, err := json.Marshal(resp)
	if err != nil {
		re.Body = err.Error()
		re.StatusCode = 400
		return re, nil
	}

	re.Body = string(jsonba)
	re.StatusCode = 200
	return re, nil
}

func main() {
	lambda.Start(handler)
}
