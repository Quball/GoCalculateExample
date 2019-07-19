# GoCalculateExample

An example service to show calculation of numbers with Golang, AWS Lambda and AWS API Gateway

## Deployment package and publish changes to AWS Lambda

`$ GOOS=linux go build cmd/main.go`

`$ zip function.zip main`

`$ aws lambda update-function-code --function-name calculateProxy --zip-file fileb://function.zip`

## Test calculate package

from the `test/` folder run
`$ go test`

## Test endpoint in Postman

POST `https://jcgpzw51yg.execute-api.eu-north-1.amazonaws.com/beta/calculateresource`

Request Body

```
{
	"number_a": 1,
	"number_b": 2,
	"operation": "SUBTRACT"
}
```

Result
```
{
    "statusCode": 200,
    "headers": {
        "Content-Type": "application/json"
    },
    "result": -1
}
```
