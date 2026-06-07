## Todo Application backend in Go

#### This is a simple Todo application backend implemented in Go with native net/http library. It provides RESTful APIs to manage todo items, including creating, retrieving, updating, and deleting todos.

API Endpoints:
- `POST /todos`: Create a new todo item.
- `GET /todos`: Retrieve all todo items.
- `GET /todos/{id}`: Retrieve a specific todo item by ID.
- `PUT /todos/{id}`: Update a specific todo item by ID.
- `DELETE /todos/{id}`: Delete a specific todo item by ID.

### Running the Application
1. Ensure you have Go installed on your machine.
2. Clone the repository and navigate to the project directory.
3. Run the application using the command:
```bash
1. make
2. ./main
```
4. The server will start on `http://localhost:3000`.