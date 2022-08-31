package csvToDB

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io"
	"log"
	"os"
)

type AuthHandler struct {
	collection *mongo.Collection
}

func NewAuthHandler(collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
	}
}

var authHandler *AuthHandler

// Connection URI
const uri = "mongodb://admin:password@localhost:27017/?maxPoolSize=20&w=majority"

func init() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

type Person struct {
	Scantime     string
	IsForeigner  string
	Ps           string
	SystemSource string
	CompanyName  string
	Name         string
	IdCard       string
	PhoneNumber  string
	Color        string
	ScanMethod   string
	Longitude    string
	Latitude     string
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	coll := client.Database("xd").Collection("shssm")

	f, err := os.Open("../../../../Downloads/scanlog_5m.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var newPerson = Person{}
		newPerson.Scantime = rec[0]
		newPerson.IsForeigner = rec[1]
		newPerson.Ps = rec[2]
		newPerson.SystemSource = rec[3]
		newPerson.CompanyName = rec[4]
		newPerson.Name = rec[5]
		newPerson.IdCard = rec[6]
		newPerson.PhoneNumber = rec[7]
		newPerson.Color = rec[8]
		newPerson.ScanMethod = rec[9]
		newPerson.Longitude = rec[10]
		newPerson.Latitude = rec[11]
		_, error := coll.InsertOne(context.TODO(), newPerson)
		if error != nil {
			panic(error)
		}
	}
}
