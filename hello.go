package main


import (
        
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}


// // main.go
// package main

// import (
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// func Handler(request Request) (Response, error) {
// 	return "Hello!", nil
// }

// func main() {
// 	// Make the handler available for Remote Procedure Call by AWS Lambda
// 	lambda.Start(Handler)
// }