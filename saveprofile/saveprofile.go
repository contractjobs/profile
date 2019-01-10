package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/contractjobs/profile/service"
)

func saveProfile(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id, err := service.NewProfileService().NewProfile(req.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(id),
	}, nil
}

func main() {
	lambda.Start(saveProfile)
}
