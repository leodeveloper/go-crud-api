# Go RESTful API with Gorilla Mux

This project demonstrates a simple RESTful API written in Go (Golang) using the Gorilla Mux router for handling HTTP requests.

## Features

- **CRUD Operations**: Supports Create, Read, Update, and Delete operations for managing educational records.
- **JSON Serialization**: Data is exchanged in JSON format between the client and the server.
- **Sample Data**: Initial sample data for educational records is provided for testing purposes.
- **Dynamic ID Generation**: IDs for new records are dynamically generated using a random number generator.

## Prerequisites

- Go (Golang) installed on your system.
- `github.com/gorilla/mux` package for routing.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/your-repo.git
    ```

2. Navigate to the project directory:

    ```bash
    cd your-repo
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

1. Start the server:

    ```bash
    go run main.go
    ```

2. The server will start running on port 8081 by default. You can access the API endpoints using tools like cURL or Postman.

## API Endpoints

- `GET /educations`: Fetch all educations.
- `GET /educations/{id}`: Fetch an education by ID.
- `POST /educations`: Create a new education.
- `PUT /educations/{id}`: Update an education by ID.
- `DELETE /educations/{id}`: Delete an education by ID.

## Sample Request (using cURL)

1. Fetch all educations:

    ```bash
    curl http://localhost:8081/educations
    ```

2. Fetch a specific education by ID:

    ```bash
    curl http://localhost:8081/educations/1
    ```

3. Create a new education:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"degree":"PhD","gpa":"3.5","student":{"firstname":"John","lastname":"Doe"}}' http://localhost:8081/educations
    ```

4. Update an existing education:

    ```bash
    curl -X PUT -H "Content-Type: application/json" -d '{"degree":"Bachelor","gpa":"3.2","student":{"firstname":"Alice","lastname":"Smith"}}' http://localhost:8081/educations/1
    ```

5. Delete an education:

    ```bash
    curl -X DELETE http://localhost:8081/educations/1
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
