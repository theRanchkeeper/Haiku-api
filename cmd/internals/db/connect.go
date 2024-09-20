package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// configuring database
type Database struct {
	dbName string
	uri    string
	client mongo.Client
}

func (d *Database) setupDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	d.dbName = os.Getenv("DB_NAME")
	d.uri = os.Getenv("MONGO_URI")

}

func GetDBClient() mongo.Client {
	return dataBase.client
}

var dataBase Database

// when calling this package in server/main.go database will be connected
func init() {

	dataBase.setupDB()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dataBase.uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	dataBase.client = *client
}
