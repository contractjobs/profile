package dynamolayer

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/contractjobs/profile/lib/persistence"
)

const (
	PROFILE = "Profile"
)

type DynamoDBLayer struct {
	service *dynamodb.DynamoDB
}

func NewDynamoDBLayerByRegion(region string) (persistence.DatabaseHandler, error) {
	log.Println("establishing dynamo session")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	log.Println("creating dynamo session")
	return &DynamoDBLayer{
		service: dynamodb.New(sess),
	}, nil
}

func NewDynamoDBLayerBySession(sess *session.Session) persistence.DatabaseHandler {
	return &DynamoDBLayer{
		service: dynamodb.New(sess),
	}
}

func (dynamoLayer *DynamoDBLayer) AddUpdateProfile(profile persistence.Profile) (string, error) {
	log.Println("in Add update Profile")
	av, err := dynamodbattribute.MarshalMap(profile)
	if err != nil {
		return "", err
	}
	_, err = dynamoLayer.service.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(PROFILE),
		Item:      av,
	})
	if err != nil {
		return "", err
	}
	log.Println("Returning profile")
	log.Println(profile.Email)
	return profile.Email, nil
}

func (dynamoLayer *DynamoDBLayer) FindAllProfiles() ([]persistence.Profile, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(PROFILE),
	}
	result, err := dynamoLayer.service.Scan(input)
	if err != nil {
		return nil, err
	}
	profile := []persistence.Profile{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &profile)
	return profile, err

}

func (dynamoLayer *DynamoDBLayer) FindProfileByEmail(name string) (persistence.Profile, error) {
	//Create the QueryInput type with the information we need to execute the query
	input := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("Email = :n"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(name),
			},
		},
		//		IndexName: aws.String(EMAIL_INDEX),
		TableName: aws.String(PROFILE),
	}
	// Execute the query
	result, err := dynamoLayer.service.Query(input)
	if err != nil {
		return persistence.Profile{}, err
	}
	//Obtain the first item from the result
	profile := persistence.Profile{}
	if len(result.Items) > 0 {
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &profile)
	} else {
		err = errors.New("No results found")
	}
	return profile, err
}
