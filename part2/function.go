package cloudfunctions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"

  // Built-in Golang packages
  "io/ioutil" // io.ReadFile
  "reflect" // get an object type

  // Import the JSON encoding package
  "encoding/json"

  // Official 'mongo-go-driver' packages
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoFields struct {
	//Key string `json:"key,omitempty"`
	ID int `json:"_id"`
	FieldStr string `json:"Field Str"`
	FieldInt int `json:"Field Int"`
	FieldBool bool `json:"Field Bool"`
}

func ListNetwork(w http.ResponseWriter, r *http.Request){

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
			log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
			log.Fatal(err)
	}

	// Project ID for this request.
	project := "my-project" // TODO: Update placeholder value.

	// Name of the region scoping this request.
	region := "my-region" // TODO: Update placeholder value.

	req := computeService.Subnetworks.List(project, region)
	if err := req.Pages(ctx, func(page *compute.SubnetworkList) error {
			for _, subnetwork := range page.Items {
					// TODO: Change code below to process each `subnetwork` resource:
					
					insert(byte(subnetwork))
					fmt.Printf("%#v\n", subnetwork)
			}
			return nil
	}); err != nil {
			log.Fatal(err)
	}
}

function insert(byte[] byteValues){

	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
	log.Fatalf("mongo.Connect() ERROR: %v", err)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		// Access a MongoDB collection through a database
	col := client.Database("JSON_docs").Collection("JSON Collection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "n")

	// Declare an empty slice for the MongoFields docs
	var docs []MongoFields
	err = json.Unmarshal(byteValues, &docs)
	// Call the InsertOne() method and pass the context and doc objects
	result, insertErr := col.InsertOne(ctx, doc)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
	} else {
		fmt.Println("InsertOne() API result:", result)
	}
}

