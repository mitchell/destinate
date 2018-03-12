package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchelljfs/destinate/backend/schema"
)

func handler() (string, error) {
	schema.Migrate()
	return "DB Succesfully AutoMigrated", nil
}

func main() {
	lambda.Start(handler)
}
