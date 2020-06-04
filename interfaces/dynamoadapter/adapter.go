package dynamoadapter

import (
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type (
	TripletStoreDynamoAdapter struct {
		Config         *aws.Config
		Service        *dynamodb.DynamoDB
		DB             *dynamo.DB
		DataSourceName string
	}
)

func NewAdapter(config *aws.Config, ds string) *TripletStoreDynamoAdapter {
	a := &Adapter{}
	a.Config = config
	a.DataSourceName = ds
	a.Service = dynamodb.New(session.New(config), a.Config)
	a.DB = dynamo.New(session.New(), a.Config)
	return a
}

func (a* TripletStoreDynamoAdapter) LoadEntity(entityid string){
	
}

func (a* TripletStoreDynamoAdapter) CreateTable(){
	
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("subject"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("property"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("objectvalue"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("subject"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("property"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(a.DataSourceName),
	}
}

func (a* TripletStoreDynamoAdapter) AddSingleItem(t EntityPropertyTriplet){


	ept, err := dynamodbattribute.MarshalMap(t)
	input := &dynamodb.PutItemInput{
		Item:ept,
		TableName:aws.String(a.DataSourceName),
	}
	_,err = svc.PutItem(input)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

}
func  (a* TripletStoreDynamoAdapter) ReadOneItem(t EntityPropertyTriplet){
	
	key, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	input := &dynamodb.GetItemInput{
		Key:key,
		TableName: aws.String(a.DataSourceName),
	}
	
	result, err := svc.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	movie := Movie{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("year: ", movie.Year)
	fmt.Println("title: ", movie.Title)

}

