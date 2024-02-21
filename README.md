<!-- # Title: Product Management API Documentation

# Overview:
This Go program implements a simple RESTful API for managing products. It allows users to retrieve all products, add new products, and retrieve a single product by its ID. The API is built using the Gin web framework.

# API ENDPOINTS


1.  GET /products
   - Description: Retrieves all products available in the system.
   - Response: Returns a JSON array containing details of all products.
   - Example:
    # Request: GET /products
                   [
                 {"id": "1", "product": "freezer", "description": "One door standing freezer", "price": 200456},
                 {"id": "2", "product": "table", "description": "a wooden table well fuenished", "price": 10676},
                 ...
               ]
     Response: Status 200 OK
      

2.  POST /addProducts
   - Description: Adds a new product to the system.
   - Request: Requires a JSON object containing details of the new product (id, product, description, price).
   - Response: Returns the details of the newly added product.
   - Example:
     # Request: POST /addProducts
              {
                "id": "5",
                "product": "chair",
                "description": "comfortable chair with armrests",
                "price": 7890
              }
     Response: Status 201 Created
               {
                 "id": "5",
                 "product": "chair",
                 "description": "comfortable chair with armrests",
                 "price": 7890
               }

3.  GET /getProduct/:id
   - Description: Retrieves a single product by its ID.
   - Request: Requires the ID of the product to be specified in the URL.
   - Response: Returns the details of the product if found, or a "Product Not Found" message otherwise.
   - Example:
     # Request: GET /getProduct/3
     Response: Status 200 OK
          
               {
                 "id": "3",
                 "product": "3D light",
                 "description": "RGB light lamb for night",
                 "price": 9893
               }
          
    # Request: GET /getProduct/10
     Response: Status 404 Not Found
     
               {
                 "message": "Product Not Found"
               }
     

# Data Structure:
- The product data is stored in an in-memory slice of product structs.
- Each product struct contains fields for ID, product name, description, and price.

# Dependencies:
- The program uses the Gin web framework for handling HTTP requests and responses.

# Usage:
- Start the server by running the main function.
- The server listens on port 9090 by default. Update the port as needed. -->

# Product Management API

This program implements a simple RESTful API for managing products. It uses the Gin framework to handle HTTP requests and responses.

## Endpoints

### `GET /products`
- Retrieves all products.
- Response: Returns a JSON array containing all products.

### `GET /products/:merchantID`
- Retrieves products associated with a specific merchant.
- Parameters:
  - `merchantID`: ID of the merchant.
- Response: Returns a JSON array containing products associated with the specified merchant ID.

### `POST /products`
- Creates a new product.
- Request Body: JSON object representing the new product.
- Response: Returns the created product with status code 201 if successful.

### `PUT /products/:id`
- Updates an existing product.
- Parameters:
  - `id`: ID of the product to be updated.
- Request Body: JSON object representing the updated product.
- Response: Returns the updated product with status code 200 if successful.

### `DELETE /products/:id`
- Deletes an existing product.
- Parameters:
  - `id`: ID of the product to be deleted.
- Response: Returns status code 204 (No Content) if successful.

## Data Model

### Product
- `ID`: Unique identifier for the product.
- `Name`: Name of the product.
- `Description`: Description of the product.
- `Price`: Price of the product.
- `CreatedAt`: Timestamp indicating when the product was created.

## Sample Usage



#  Dependencies

## github.com/gin-gonic/gin
Gin web framework used for handling HTTP requests and responses

```bash
# Start the server
go run server.go




