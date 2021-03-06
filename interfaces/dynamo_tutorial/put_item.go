package main 

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main(){

	type MovieInfo struct {
		Plot string 	`json:"plot"`
		Rating float64	`json:"rating"`
	}
	type Movie struct {
		Year int 	`json:"year"`
		Title string 	`json:"title"`
		Info MovieInfo `json:"info"`
	}

	movie := Movie{
		Year:2015,
		Title:"Big new movie",
		Info:MovieInfo{
			Plot:"Nothing happend here",
			Rating:0.0,
		},
	}
	config := &aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(movie)
	input := &dynamodb.PutItemInput{
		Item:av,
		TableName:aws.String("Movies"),
	}
	_,err = svc.PutItem(input)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("We have inserted a new item")

}