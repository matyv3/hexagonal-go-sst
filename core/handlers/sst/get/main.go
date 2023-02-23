package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matyv3/hexagonal-go-sst/core/repositories"
	"github.com/matyv3/hexagonal-go-sst/core/services"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	repository := repositories.CreateTODORepository()
	service := services.CreateTODOService(repository)

	result, err := service.GetTODOs()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 201,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
