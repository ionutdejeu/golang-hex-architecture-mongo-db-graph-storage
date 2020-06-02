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

}

