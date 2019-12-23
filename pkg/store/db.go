package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB is the mongo DB instance
type DB struct {
	Name       string
	URL        string
	Collection string
}

//NewDB returns a new DB object
func NewDB(url, name, collection string) *DB {
	return &DB{
		Name:       name,
		URL:        url,
		Collection: collection,
	}
}

func (db *DB) connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(db.URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("error conencting to mongo db. %+v", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("error pinging to mongo db. %+v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil

}

//InsertOHLCResult single ohlc result document into the database
func (db *DB) InsertOHLCResult(result *Result) error {
	client, err := db.connect()
	defer client.Disconnect(context.TODO())
	if err != nil {
		return fmt.Errorf("error inserting document to db. %+v", err)
	}
	collection := client.Database(db.Name).Collection(db.Collection)

	insertResult, err := collection.InsertOne(context.TODO(), &result)
	if err != nil {
		return fmt.Errorf("error inserting document to db. %+v", err)
	}

	fmt.Println("Inserted a OHLC document: ", insertResult.InsertedID)

	return nil
}

//GetOHLCResult returns the last ohlc analysis results from the db
func (db *DB) GetOHLCResult() (*Result, error) {
	client, err := db.connect()
	defer client.Disconnect(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error findng document in db. %+v", err)
	}
	collection := client.Database(db.Name).Collection(db.Collection)

	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	log.Println(count)
	// //var filter bson.D
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var result Result
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error finding document in db. %+v", err)
	}
	//cur.Decode(&results)
	defer cur.Close(ctx)
	var tc int64
	for cur.Next(ctx) {
		if tc == count-1 {
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			break
		}

		tc = tc + 1
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//var result Result
	// err = collection.FindOne(context.TODO(), bson.M{"$natural": -1}).Decode(&result)
	// if err != nil {
	// 	return nil, fmt.Errorf("error finding document in db. %+v", err)
	// }

	// dbSize, err := collection.CountDocuments()
	// if err != nil {

	// }

	// err = c.Find(nil).skip(dbSize-1).One(&myData)
	// if err != nil {
	//         return err
	// }

	//cur, err := collection.Find(context.TODO(), nil).Select(bson.M{"numbers": 1}).Sort("-$natural").Limit(1).One(&numbs)

	return &result, nil

}
