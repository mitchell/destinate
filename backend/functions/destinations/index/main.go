package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchelljfs/destinate/backend/schema"
)

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var re events.APIGatewayProxyResponse
	var u schema.User
	uid := r.PathParameters["id"]
	query := "food shops activities venice, ca"

	if uid != "" {
		if err := u.Find(uid); err != nil {
			re.Body = err.Error()
			re.StatusCode = 400
			return re, nil
		}

		if u.Location != "" && u.Scores.RawMessage != nil {
			val, err := u.Scores.Value()
			if err != nil {
				re.Body = err.Error()
				re.StatusCode = 400
				return re, nil
			}

			jmap := map[string]float64{}

			if val != nil {
				if err := json.Unmarshal(val.([]byte), &jmap); err != nil {
					re.Body = err.Error()
					re.StatusCode = 400
					return re, nil
				}
			}

			scores := schema.RankByWordCount(jmap)

			query = fmt.Sprintf("%s %s %s %s", scores[0].Key, scores[1].Key, scores[2].Key, u.Location)
		} else if u.Location != "" {
			query = fmt.Sprintf("food shops activities %s", u.Location)
		} else if u.Scores.RawMessage != nil {
			val, err := u.Scores.Value()
			if err != nil {
				re.Body = err.Error()
				re.StatusCode = 400
				return re, nil
			}

			jmap := map[string]float64{}

			if val != nil {
				if err := json.Unmarshal(val.([]byte), &jmap); err != nil {
					re.Body = err.Error()
					re.StatusCode = 400
					return re, nil
				}
			}

			scores := schema.RankByWordCount(jmap)

			query = fmt.Sprintf("%s %s %s venice, ca", scores[0].Key, scores[1].Key, scores[2].Key)
		}
	}
	log.Printf("query: %v", query)

	resp, err := schema.AllDestinations(ctx, query)
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
