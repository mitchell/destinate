package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"googlemaps.github.io/maps"
)

func handler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCWhNQ3cPHw7qFiiRTucVd61W4ZSzeUyTI"))
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}

	tsr := &maps.TextSearchRequest{
		Query: "food near venice",
	}

	resp, err := c.TextSearch(ctx, tsr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}

	// ppr := &maps.PlacePhotoRequest{
	// 	PhotoReference: resp.Results[0].Photos[0].PhotoReference,
	// }
	// imgResp, err := c.PlacePhoto(ctx, ppr)
	// img, err := &imgResp.Image()
	// log.Printf("image: %v", img)

	jsonba, err := json.Marshal(resp.Results)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonba),
		StatusCode: 200,
	}, err
}

func main() {
	lambda.Start(handler)
}
