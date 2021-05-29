package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	apigatewaylambdagoexample "github.com/DanilloDeSouza/api-gateway-lambda-go-example/api-gateway-lamba-go-example"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	config := apigatewaylambdagoexample.NewConfig()
	logger, err := apigatewaylambdagoexample.NewLogger(config.LogLevel)
	uri := req.Path

	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sqs, err := apigatewaylambdagoexample.NewSqs(config.AWSRegion, config.SQSEndpoint)
	if err != nil {
		apigatewaylambdagoexample.LogError(ctx, logger, fmt.Sprintf("could not init sqs: %s", err), uri)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	service := apigatewaylambdagoexample.NewService(*sqs)

	jsonString, jsonErr := json.Marshal(req.QueryStringParameters)
	if jsonErr != nil {
		apigatewaylambdagoexample.LogError(ctx, logger, fmt.Sprintf("could not marshal querystring from GET: %s, body: %s", jsonErr, req.QueryStringParameters), uri)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	createErr := service.CreateMessage(ctx, config.SQSEndpoint, jsonString)

	if createErr != nil {
		apigatewaylambdagoexample.LogError(ctx, logger, fmt.Sprintf("could not create a sqs message: %s", createErr), uri)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	config := apigatewaylambdagoexample.NewConfig()
	logger, err := apigatewaylambdagoexample.NewLogger(config.LogLevel)

	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	lambda.Start(handle)
}
