

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



#  Dependencies

## github.com/gin-gonic/gin
Gin web framework used for handling HTTP requests and responses




#     Entity-Relationship Diagram (ERD)


            +---------------------+
            |       Merchant      |
            +---------------------+
            | merchant_id (PK)    |
            | name                |
            | email               |
            | ...                 |
            +---------------------+
                    |
                    |
                    |
                    |
            +---------------------+
            |       Product       |
            +---------------------+
            | product_id (PK)     |
            | name                |
            | description         |
            | price               |
            | created_at          |
            | merchant_id (FK)    |
            +---------------------+


#     Database Design Considerations:


# Merchant Table:
- This table stores information about each merchant, such as their `name`, `email`, etc.
- The `merchant_id` serves as the primary key (PK) to uniquely identify each merchant.
- Additional attributes like contact information, address, etc., can be added based on requirements.
- To optimize performance, proper indexing should be applied to columns used frequently in search or filtering operations.

# Product Table:

- This table stores information about each product, including its `name`, `description`, `price`, and the `merchant` it belongs to.
- The product_id serves as the primary key (PK) to uniquely identify each product.
- The `merchant_id` is a foreign key (FK) referencing the `merchant_id` in the Merchant table, establishing a relationship between products and merchants.
Indexing should be applied to the `merchant_id` column to facilitate efficient retrieval of products by merchant.
 
#  Performance Optimization Strategies:


# Indexing:
- Indexes should be created on columns frequently used in search, filtering, and joining operations, such as `merchant_id`.
- Proper indexing can significantly improve query performance, especially when dealing with large datasets.

# Database Sharding:

- For extremely large datasets, consider database sharding to horizontally partition the data across multiple database instances.
- Sharding based on `merchant ID` or other relevant criteria can distribute the workload and improve scalability and performance.

# Caching:

- Implement caching mechanisms to store frequently accessed data in memory, reducing database load and query response times.
- Utilize caching solutions like Redis or Memcached to cache frequently accessed product data.

#        Choice of Database

When selecting a database for this problem, consider the following factors:

# Scalability:

- Choose a database that can scale horizontally to handle a large number of merchants and products efficiently.
- NoSQL databases like MongoDB or Cassandra are often preferred for their ability to scale out across multiple nodes.


# Performance:

- Look for databases optimized for read-heavy workloads, as retrieving product data for display is a common operation.
- Consider databases with robust indexing and caching capabilities to optimize query performance.

# Data Consistency:

- Ensure data consistency and ACID compliance, especially in scenarios involving transactions or updates to product information.
- Traditional relational databases like PostgreSQL or MySQL are well-suited for maintaining data consistency.

# Flexibility:

- Choose a database that offers flexibility in schema design to accommodate evolving business requirements.
- NoSQL databases provide schema flexibility, making them suitable for agile development and handling diverse data structures.

Based on these considerations, a combination of a scalable NoSQL database for product storage and a relational database for merchant information could be an effective solution. However, the specific choice of database depends on factors such as the application's requirements, data volume, expected growth, and available resources.


## Sample Usage



```bash
# Start the server
go run server.go