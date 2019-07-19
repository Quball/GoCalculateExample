package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Calculate struct {
	NumberA string `json:"number_a"`
	NumberB string `json:"number_b"`
}

func HandleRequest(ctx context.Context, name Calculate) (string, error) {
	return fmt.Sprintf("Calculate number %s with %s", name.NumberA, name.NumberB), nil
}

func main() {
	lambda.Start(HandleRequest)
}
