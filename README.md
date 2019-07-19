# GoCalculateExample

An example service to show calculation of numbers with Golang, AWS Lambda and AWS API Gateway

## Deployment package and publish changes to AWS Lambda

`$ GOOS=linux go build cmd/main.go`

`$ zip function.zip main`

`$ aws lambda update-function-code --function-name calculateProxy --zip-file fileb://function.zip`

## Test endpoint in Postman

POST `https://jcgpzw51yg.execute-api.eu-north-1.amazonaws.com/beta/calculateresource`

Request Body

```
{
	"number_a": "13",
	"number_b": "42"
}
```

Result: "Calculate number 13 with 42"
