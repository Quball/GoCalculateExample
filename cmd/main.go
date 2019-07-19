package main

import (
	"context"
	"errors"

	"./calculate"

	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	NumberA   int    `json:"number_a"`
	NumberB   int    `json:"number_b"`
	Operation string `json:"operation"`
}

type response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Result     float32           `json:"result"`
}

// HandleRequest returns operation for calculate or error
func HandleRequest(ctx context.Context, name request) (response, error) {
	result, err := calculate.Calculate(name.NumberA, name.NumberB, name.Operation)
	if err != nil {
		return response{},
			errors.New(`Math operation is not recognized`)
	}
	return response{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Result:     result,
		},
		nil
}

func main() {
	lambda.Start(HandleRequest)
}
