package main

import (
	"context"
	"errors"
	"fmt"

	"./calculate"

	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	NumberA   int    `json:"number_a"`
	NumberB   int    `json:"number_b"`
	Operation string `json:"operation"`
}

// HandleRequest returns operation for calculate or error
func HandleRequest(ctx context.Context, name request) (string, error) {
	result, err := calculate.Calculate(name.NumberA, name.NumberB, name.Operation)
	if err != nil {
		return "", errors.New(`Math operation is not recognized`)
	}
	return fmt.Sprintf("%s (%d, %d) = %f", name.Operation, name.NumberA, name.NumberB, result), nil
}

func main() {
	lambda.Start(HandleRequest)
}
