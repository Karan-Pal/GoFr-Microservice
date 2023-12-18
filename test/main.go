package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCRUDOperations(t *testing.T) {
	app := fiber.New()
	// Adding a cors for put and delete request	
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		return c.Next()
	})
	app.Get("/tiles", getTiles)
	app.Get("/tiles/:id", getTile)
	app.Post("/tiles", createTile)
	app.Put("/tiles/:id", updateTile)
	app.Delete("/tiles/:id", deleteTile)
	
	t.Run("CreateTile", func(t *testing.T) {
		// Define a sample tile for testing
		tile := Tile{
			Name:         "Kajaria",
			ID:			  "1",
			Stock:		  "10",
		}

		// Convert tile to JSON
		tileJSON, err := json.Marshal(tile)
		assert.NoError(t, err)

		// Create a POST request with the tile JSON
		req := httptest.NewRequest(http.MethodPost, "/tiless", bytes.NewReader(tileJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		// Assert the status code is 200 OK
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("GetTiles", func(t *testing.T) {
		// Create a GET request to retrieve all tiles
		req := httptest.NewRequest(http.MethodGet, "/tiles", nil)

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		// Assert the status code is 200 OK
		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

	
	t.Run("UpdateTile", func(t *testing.T) {
		// Define a sample tile for testing
		updatedTile := Tile{
			Name:         "Somany",
			ID:			  "2",
			Stock:		  "100",
			
		}

		// Convert updated tile to JSON
		updatedTileJSON, err := json.Marshal(updatedTile)
		assert.NoError(t, err)

		// Create a PUT request with the updated tile JSON
		req := httptest.NewRequest(http.MethodPut, "/tiles/657ad262660b9c2342474c33", bytes.NewReader(updatedTileJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("DeleteTile", func(t *testing.T) {
		// Create a DELETE request to delete a tile
		req := httptest.NewRequest(http.MethodDelete, "/tiles/657be72ea13c844c358f83b2", nil)

		// Create a response recorder to record the response
		res, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

}


