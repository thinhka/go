// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type item struct {
// 	Id        int        `json: "id"`
// 	ShortDes  string     `json: "short_des"`
// 	CreatedAt *time.Time `json: "created_at"`
// }

// type Recipe struct {
// 	Name              string
// 	Ingredients       []string
// 	PrepTimeInMinutes int `json:"prepTimeInMinutes" bson:"prepTimeInMinutes"`
// }

// func main() {
// 	// fmt.Println("h")

// 	// now := time.Now().UTC()
// 	// startItem := item{
// 	// 	Id:        1,
// 	// 	ShortDes:  "new item",
// 	// 	CreatedAt: &now,
// 	// }

// 	// jsonData, err := json.Marshal(startItem)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// fmt.Println(string(jsonData))

// 	// r := gin.Default()
// 	// r.GET("/ping", func(c *gin.Context) {
// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"message": startItem,
// 	// 	})
// 	// })
// 	// r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

// 	var mongoUri = "mongodb+srv://thinhka:MOk6OXn0vxRp1sLF@thinhka.rtzox.mongodb.net/?retryWrites=true&w=majority&appName=thinhka"

// 	// CONNECT TO YOUR ATLAS CLUSTER:
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		mongoUri,
// 	))

// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	err = client.Ping(ctx, nil)

// 	if err != nil {
// 		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
// 		panic(err)
// 	}

// 	fmt.Println("Connected to MongoDB!\n")

// 	// Provide the name of the database and collection you want to use.
// 	// If they don't already exist, the driver and Atlas will create them
// 	// automatically when you first write data.
// 	var dbName = "myDatabase"
// 	var collectionName = "recipes"
// 	collection := client.Database(dbName).Collection(collectionName)

// 	/*
// 	 * *** FIND DOCUMENTS ***
// 	 *
// 	 * Now that we have data in Atlas, we can read it. To retrieve all of
// 	 * the data in a collection, we create a filter for recipes that take
// 	 * less than 45 minutes to prepare and sort by name (ascending)
// 	 */

// 	var filter = bson.M{"prepTimeInMinutes": bson.M{"$lt": 45}}
// 	options := options.Find()

// 	// Sort by `name` field ascending
// 	options.SetSort(bson.D{{"name", 1}})

// 	cursor, err := collection.Find(context.TODO(), filter, options)
// 	if err != nil {
// 		fmt.Println("Something went wrong trying to find the documents:")
// 		panic(err)
// 	}

// 	defer func() {
// 		cursor.Close(context.Background())
// 	}()

// 	// Loop through the returned recipes
// 	for cursor.Next(ctx) {
// 		recipe := Recipe{}
// 		err := cursor.Decode(&recipe)

// 		// If there is an error decoding the cursor into a Recipe
// 		if err != nil {
// 			fmt.Println("cursor.Next() error:")
// 			panic(err)
// 		} else {
// 			fmt.Println(recipe.Name, "has", len(recipe.Ingredients), "ingredients, and takes", recipe.PrepTimeInMinutes, "minutes to make.\n")
// 		}
// 	}

// 	// We can also find a single document. Let's find the first document
// 	// that has the string "potato" as an ingredient
// 	var result Recipe
// 	var myFilter = bson.D{{"ingredients", "potato"}}
// 	e := collection.FindOne(context.TODO(), myFilter).Decode(&result)
// 	if e != nil {
// 		fmt.Println("Something went wrong trying to find one document:")
// 		panic(e)
// 	}
// 	fmt.Println("Found a document with the ingredient potato", result, "\n")

// }
