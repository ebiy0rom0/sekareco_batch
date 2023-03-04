package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Run)
}

func Run(ctx context.Context) error {
	return nil
}
