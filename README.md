# GoFr-Microservice
The Go CRUD Application is a simple yet robust project built using the Fiber web framework and MongoDB database. 
It facilitates basic CRUD operations (Create, Read, Update, Delete) for managing tile records. Fork, customize, 
and extend to meet your specific needs with the flexibility of Go and Fiber.

This is a microservice made using Golang framework for my family business Yash Marble & Tiles.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [MongoDB](https://www.mongodb.com/docs/manual/administration/install-community/)
- [PostMan](https://www.postman.com/downloads/)

 ### Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/Karan-Pal/GoFr-Microservice.git
   
   cd GoFr-Microservice
2. Download Dependencies:

   ```bash
   go mod download

3. Verify Dependencies:

    ```bash
   cat go.sum
4. Run Project
    
    ```bash
    go run main.go

5. Open Server
    
    ```bash
    http://localhost:8080/tiles

 #### API Endpoints

The application provides RESTful API endpoints for CRUD operations on tiles records.

### Base URL

The base URL for all endpoints is:

### List of Endpoints

#### 1. **Get All Tiles**

- **Endpoint:**
  - `GET /tiles`

- **Description:**
  - Retrieves a list of all tiles.

 ##### 2. **Get Tilet by ID**

- **Endpoint:**
  - `GET /tiles/:id`
- **Parameters:**
  - `id`: The unique identifier for the tile.

- **Description:**
  - Retrieves information about a specific tile based on the provided ID.
 
 #### 3. **Create a New Tile**

- **Endpoint:**
  - `POST /tiles`
- **Request Body:**
  ```json
  
 **Description:**
  - Creates a new tile record.

#### 4. **Update Tile by ID**

- **Endpoint:**
  - `PUT /tiles/:id`
- **Parameters:**
  - `id`: The unique identifier for the tile.

- **Request Body:**
  ```json

- **Description:**
  - Updates information about a specific tile based on the provided ID.
 
#### 5. **Delete Tile by ID**

- **Endpoint:**
  - `DELETE /tiles/:id`

- **Parameters:**
  - `id`: The unique identifier for the tile.

- **Description:**
  - Deletes a specific tiles record based on the provided ID.

- **Example Response:**
  ```json
  "Tile deleted successfully"

### Running Tests

- Use the following command to run the tests:

  ```bash
  go test
