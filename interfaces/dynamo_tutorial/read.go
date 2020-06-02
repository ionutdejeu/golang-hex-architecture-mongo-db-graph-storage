package main 

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main(){
	
	type Movie struct {
		Year  int         `json:"year"`
		Title string      `json:"title"`
		Info  interface{} `json:"info,omitempty"`
	}

	config := &aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess :=  session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)
	movieKey := Movie{
		Year:2015,
		Title:"Big new movie",
	}

	key, err := dynamodbattribute.MarshalMap(movieKey)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	input := &dynamodb.GetItemInput{
		Key:key,
		TableName: aws.String("Movies"),
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

	infoMap := movie.Info.(map[string]interface{})
	for k, v := range infoMap {
		switch vv := v.(type) {
		case string, float64:
			fmt.Println(k, ": ", vv)
		case []interface{}:
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}