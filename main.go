// main.go

package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"gofr.dev/pkg/gofr"
	"log"
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

// Initialize initializes the inventory service.
func (s *InventoryService) Initialize() {
	s.Tiles = make(map[string]Tile)
}

// CreateTileHandler handles the creation of a new tile.
func (s *InventoryService) CreateTileHandler(ctx *gofr.Context) (interface{}, error) {
	var tile Tile
	if err := ctx.WebSocketConnection.ReadJSON(&tile); err != nil {
		return nil, err
	}

	// Generate a unique ID (you might want to use a more robust method)
	tile.ID = fmt.Sprintf("%d", len(s.Tiles)+1)

	s.Tiles[tile.ID] = tile
	return tile, nil
}

// UpdateStockHandler handles the update of tile stock.
func (s *InventoryService) UpdateStockHandler(ctx *gofr.Context) (interface{}, error) {
	tileID := ctx.PathParam("id")

	// Assuming a positive or negative integer is sent in the request body for stock update
	var stockUpdate int
	if err := ctx.WebSocketConnection.ReadJSON(&stockUpdate); err != nil {
		return nil, err
	}

	tile, found := s.Tiles[tileID]
	if !found {
		return nil, fmt.Errorf("Tile %v not found)", tileID)
	}

	// Update stock
	tile.Stock += stockUpdate
	s.Tiles[tileID] = tile

	return tile, nil
}

// ReadTilesHandler handles reading all available tiles.
func (s *InventoryService) ReadTilesHandler(ctx *gofr.Context) (interface{}, error) {
	tiles := make([]Tile, 0, len(s.Tiles))
	for _, tile := range s.Tiles {
		tiles = append(tiles, tile)
	}
	return tiles, nil
}

// DeleteTileHandler handles deleting finished stock from the database.
func (s *InventoryService) DeleteTileHandler(ctx *gofr.Context) (interface{}, error) {
	tileID := ctx.PathParam("id")

	// Check if the tile exists
	_, found := s.Tiles[tileID]
	if !found {
		return nil, fmt.Errorf("Tile %v not found)", tileID)
	}

	// Delete the tile
	delete(s.Tiles, tileID)

	return map[string]string{"message": "Tile deleted successfully"}, nil
}

func main() {
	app := gofr.New()

	// Initialize the InventoryService
	inventoryService := &InventoryService{}
	inventoryService.Initialize()

	// Register handlers
	app.POST("/tiles", inventoryService.CreateTileHandler)
	app.PUT("/tiles/:id/stock", inventoryService.UpdateStockHandler)
	app.GET("/tiles", inventoryService.ReadTilesHandler)
	app.DELETE("/tiles/:id", inventoryService.DeleteTileHandler)

	// Specify the port for the server to listen on
	/*port := 8080
	app.Start()

	// Start the server
	fmt.Printf("Server is running on port %d...\n", port)
	*/

	cluster := gocql.NewCluster("localhost")
	cluster.Port = 2011
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	app.Start()
}
