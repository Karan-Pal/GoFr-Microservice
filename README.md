# GoFr-Microservice

This is a microservice made using Golang framework for my family business Yash Marble & Tiles. It has following CRUD Operations:
CREATE
READ
UPDATE
DELETE

To run the service, run the command, go run main.go

These are the following four commands for CRUD Operations:
To create a tile, curl -X POST -d '{"name":"TileName", "stock":10}' http://localhost:8080/tiles
To read all tiles, curl http://localhost:8080/tiles
To update the tiles, curl -X PUT -d '2' http://localhost:8080/tiles/TILE_ID/stock (# Replace TILE_ID with the actual ID of the tile you want to update)
To delete the tile, curl -X DELETE http://localhost:8080/tiles/TILE_ID (# Replace TILE_ID with the actual ID of the tile you want to delete)

Make sure to replace TILE_ID with the actual ID of the tile you want to update or delete.

 
