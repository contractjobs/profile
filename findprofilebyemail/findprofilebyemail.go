package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/contractjobs/profile/service"
)

func findProfileByEmail(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	email := req.PathParameters["email"]
	log.Println(email)
	resp, err := service.NewProfileService().FindProfileByEmail(email)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	log.Println("sending success response")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: resp,
	}, nil
}

func main() {
	log.Println("calling findprofilebyemail")
	lambda.Start(findProfileByEmail)
}
