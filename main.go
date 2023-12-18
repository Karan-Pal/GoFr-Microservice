// main.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Tile represents a tile in the inventory.
type Tile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

// InventoryService represents the microservice for inventory management.
type InventoryService struct {
	Tiles map[string]Tile
}

var client *mongo.Client

// Connect to MongoDB
func init() {
	// Set up client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create a MongoDB client
	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// Function to create a new student
func createTile(c *fiber.Ctx) error {
	var tile Tile
	if err := c.BodyParser(&tile); err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("tiles")
	result, err := collection.InsertOne(context.TODO(), tile)
	if err != nil {
		return err
	}

	return c.JSON(result)
}


// Function to get all tiles
func getTiles(c *fiber.Ctx) error {
	var tiles []Tile
	collection := client.Database("testdb").Collection("tiles")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var tile Tile
		cursor.Decode(&tile)
		tiles = append(tiles, tile)
	}
	return c.JSON(tiles)
}

// Function to get a single tile by ID
func getTile(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var tile Tile
	collection := client.Database("testdb").Collection("tiles")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&tile)
	if err != nil {
		return err
	}
	return c.JSON(tile)
}

// Function to update a tile by ID
func updateTile(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var updatedTile Tile
	if err := c.BodyParser(&updatedTile); err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("tiles")
	filter := bson.M{"_id": objectID}
	update := bson.D{{"$set", updatedTile}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return c.JSON("Tile updated successfully")
}

// Function to delete a tile by ID
func deleteTile(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := client.Database("testdb").Collection("tiles")
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return c.JSON("Tile deleted successfully")
}

func main() {
	app := fiber.New()

	// Define CRUD endpoints for tiles
	app.Get("/tiles", getTiles)
	app.Get("/tiles/:id", getTile)
	app.Post("/tiles", createTile)
	app.Put("/tiles/:id", updateTile)
	app.Delete("/tiles/:id", deleteTile)

	// Start the server
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

