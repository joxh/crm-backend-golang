# Customer Relationship Management Service (Golang)

Author: [Joshua Schoenfield](https://github.com/joxh)
Code: [GitHub](https://github.com/joxh/crm-backend-golang)

## Description

This application is a RESTful API service designed to manage customer information. Built with Go and leveraging the Gorilla Mux router, it offers a simple yet powerful platform for performing CRUD (Create, Read, Update, Delete) operations on customer data.

### Key Features and Functionalities

- **CRUD Operations**: Supports creating, reading, updating, and deleting customer records.
- **In-Memory Database**: Utilizes a mock database implemented as a Go map for quick and efficient data storage and retrieval.
- **Validation**: Includes basic validation for customer ID during operations, ensuring data integrity and consistency.
- **Content-Type Handling**: Responds with JSON for easy integration with front-end technologies or other services.

## Getting Started

### Prerequisites

- Go (version 1.15 or later recommended)
- Gorilla Mux (a powerful HTTP router and URL matcher for building Go web servers)

### Installation

1. **Install Go**: Follow the official Go installation guide at https://golang.org/doc/install.
2. **Get Gorilla Mux**: Run `go get -u github.com/gorilla/mux` to install Gorilla Mux.

### Running the Application

1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Run `go run main.go` to start the server. By default, the server will listen on port 3000. 

> If you wish to run on a different port, that can be controlled by setting the `CRM_PORT` environment variable. For example, with the command `CRM_PORT=3001 go run main.go`, the server will listen on port 3001.

### API Endpoints and Responses

#### Add a New Customer

- **Endpoint**: `POST /customers`
- **Description**: Creates a new customer record in the database.
- **Request Body Example**:
  ```json
  {
      "id": 4,
      "name": "New Customer",
      "role": "User",
      "email": "new.customer@example.com",
      "phone": "987-654-3210",
      "contacted": false
  }
  ```
- **Response**: Returns the entire customer table including the newly added customer if the operation is successful. If the customer ID already exists, it returns an error message.

#### Get All Customers

- **Endpoint**: `GET /customers`
- **Description**: Retrieves a list of all customers in the database.
- **Response**: Returns a JSON object containing all customer records.

#### Get Single Customer

- **Endpoint**: `GET /customers/{id}`
- **Description**: Retrieves details of a single customer by their ID.
- **Response**: If the customer exists, returns a JSON object containing the customer's details. If the customer with the specified ID does not exist, returns an error message.

#### Update Existing Customer

- **Endpoint**: `PUT /customers/{id}`
- **Description**: Updates the details of an existing customer.
- **Request Body Example**:
  ```json
  {
      "id": 1,
      "name": "Updated Customer",
      "role": "Admin",
      "email": "update.customer@example.com",
      "phone": "987-654-3210",
      "contacted": true
  }
  ```
- **Response**: If the update is successful, returns a JSON object containing all customer records. If the specified ID does not exist in the database, returns an error message.

#### Delete Customer

- **Endpoint**: `DELETE /customers/{id}`
- **Description**: Deletes a customer from the database based on their ID.
- **Response**: If the deletion is successful, returns the updated customer table without the deleted record. If the specified ID does not exist, returns an error message.



### Example Calls

#### Add a New Customer

```bash
curl -X POST http://localhost:3000/customers \
-H 'Content-Type: application/json' \
-d '{
    "id": 4,
    "name": "New Customer",
    "role": "User",
    "email": "new.customer@example.com",
    "phone": "987-654-3210",
    "contacted": false
}'
```

#### Get All Customers

```bash
curl -X GET http://localhost:3000/customers
```

#### Get Single Customer

```bash
curl -X GET http://localhost:3000/customers/1
```

#### Update Existing Customer

```bash
curl -X PUT http://localhost:3000/customers/1 \
-H 'Content-Type: application/json' \
-d '{
    "id": 1,
    "name": "Updated Customer",
    "role": "Admin",
    "email": "update.customer@example.com",
    "phone": "987-654-3210",
    "contacted": true
}'
```

#### Delete Customer

```bash
curl -X DELETE http://localhost:3000/customers/1
```
