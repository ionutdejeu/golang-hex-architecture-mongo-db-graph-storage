package mongostorage

import (
	"context"
	"time"
	"log"
	"math/rand"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type TripletRecord struct {
	Subject string
	Predicate string
	Object string
}

func main(){
	defer timeTrack(time.Now(), "Mongo DB")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Connect(ctx)
	if err != nil{

	}
	collection := client.Database("graphdb").Collection("testcollection")
	collection.InsertOne(ctx, bson.M{"PrimaryKey": "Post1", "value": 3.14159})
	//id := res.InsertedID
	triplet :=  createIt();
	log.Printf("%s",triplet)
	
	 
}

func generateTriplets(n int) { 
	for i:=0;i<=n;i++{
		createRandomEntity()
	}
	return nil;
}
func insertOne
func createRandomEntity(string id) []TripletRecord {
	s []TripletRecord

	append(s,
		createTriplet(id,"id",id),
		createTriplet(id,"name","Matt")
		createTriplet(id,"friend","John")
	)
	log.Printf("%s",s)

}
func createTriplet(subject string,predicate string, Object: string) TripletRecord { 
	return TripletRecord{Subject:"Subject",Predicate:"Predicated",Object:"Ojbect"}
}

func timeTrack(start time.Time, name string){
	elapsed := time.Since(start)
	log.Printf("%s took %s",name,elapsed)
}
