// main.go

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	DB    *sql.DB
}

func (s *InventoryService) Initialize() {
	s.Tiles = make(map[string]Tile)

	// Initialize SQLite3 database
	db, err := sql.Open("sqlite3", "./tiles.db")
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db

	// Create the 'tiles' table if not exists
	createTable := `
CREATE TABLE IF NOT EXISTS tiles (
    id TEXT PRIMARY KEY,
    name TEXT,
    stock INTEGER
);
`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTileHandler handles the creation of a new tile.
func (s *InventoryService) CreateTileHandler(ctx *gofr.Context) (interface{}, error) {
	var tile Tile
	if err := ctx.BodyParser(&tile); err != nil {
		return nil, err
	}

	// Generate a unique ID
	tile.ID = fmt.Sprintf("%d", len(s.Tiles)+1)

	// Insert into SQLite3 database
	_, err := s.DB.Exec("INSERT INTO tiles (id, name, stock) VALUES (?, ?, ?)", tile.ID, tile.Name, tile.Stock)
	if err != nil {
		return nil, err
	}

	s.Tiles[tile.ID] = tile
	return tile, nil
}

// UpdateStockHandler handles the update of tile stock.
func (s *InventoryService) UpdateStockHandler(ctx *gofr.Context) (interface{}, error) {
	tileID := ctx.PathParam("id")

	// Assuming a positive or negative integer is sent in the request body for stock update
	var stockUpdate int
	if err := ctx.BodyParser(&stockUpdate); err != nil {
		return nil, err
	}

	tile, found := s.Tiles[tileID]
	if !found {
		return nil, fmt.Errorf("tile %v not found", tileID)
	}

	// Update stock
	tile.Stock += stockUpdate
	s.Tiles[tileID] = tile

	// Update SQLite3 database
	_, err := s.DB.Exec("UPDATE tiles SET stock=? WHERE id=?", tile.Stock, tileID)
	if err != nil {
		return nil, err
	}

	return tile, nil
}

// ReadTilesHandler handles reading all available tiles.
func (s *InventoryService) ReadTilesHandler(ctx *gofr.Context) (interface{}, error) {
	tiles := make([]Tile, 0, len(s.Tiles))
	rows, err := s.DB.Query("SELECT * FROM tiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tile Tile
		err := rows.Scan(&tile.ID, &tile.Name, &tile.Stock)
		if err != nil {
			return nil, err
		}
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
		return nil, fmt.Errorf("tile %v not found", tileID)
	}

	// Delete from SQLite3 database
	_, err := s.DB.Exec("DELETE FROM tiles WHERE id=?", tileID)
	if err != nil {
		return nil, err
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
	port := 8080
	app.Start()

	// Start the server
	fmt.Printf("Server is running on port %d...\n", port)
}
