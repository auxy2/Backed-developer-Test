Title: Product Management API Documentation

Overview:
This Go program implements a simple RESTful API for managing products. It allows users to retrieve all products, add new products, and retrieve a single product by its ID. The API is built using the Gin web framework.

API Endpoints:
1. GET /products
   - Description: Retrieves all products available in the system.
   - Response: Returns a JSON array containing details of all products.
   - Example:
     Request: GET /products
     Response: Status 200 OK
      

2. POST /addProducts
   - Description: Adds a new product to the system.
   - Request: Requires a JSON object containing details of the new product (id, product, description, price).
   - Response: Returns the details of the newly added product.
   - Example:
     Request: POST /addProducts
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

3. GET /getProduct/:id
   - Description: Retrieves a single product by its ID.
   - Request: Requires the ID of the product to be specified in the URL.
   - Response: Returns the details of the product if found, or a "Product Not Found" message otherwise.
   - Example:
     Request: GET /getProduct/3
     Response: Status 200 OK
               {
                 "id": "3",
                 "product": "3D light",
                 "description": "RGB light lamb for night",
                 "price": 9893
               }
             Request: GET /getProduct/10
     Response: Status 404 Not Found
               {
                 "message": "Product Not Found"
               }

Data Structure:
- The product data is stored in an in-memory slice of product structs.
- Each product struct contains fields for ID, product name, description, and price.

Dependencies:
- The program uses the Gin web framework for handling HTTP requests and responses.

Usage:
- Start the server by running the main function.
- The server listens on port 9090 by default. Update the port as needed.
