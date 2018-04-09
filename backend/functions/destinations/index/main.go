package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"googlemaps.github.io/maps"
)

func handler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCWhNQ3cPHw7qFiiRTucVd61W4ZSzeUyTI"))
	log.Printf("c: %v", c)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	tsr := &maps.TextSearchRequest{
		Query: "restaurant",
	}
	resp, err := c.TextSearch(ctx, tsr)
	log.Printf("resp: %v", resp)
	jsonba, err := json.Marshal(resp.Results)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonba),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
