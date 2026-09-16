# Go Event Registration API

A RESTful API built in Go for creating and managing events, with user authentication and event registration. Users can sign up, log in, create events, and register or cancel their registration for events created by others.

## Features

- **User authentication** — signup and login with JWT-based sessions
- **Password security** — passwords hashed with bcrypt before storage
- **Event management** — full CRUD (create, read, update, delete) for events
- **Event registration** — authenticated users can register or cancel registration for events
- **Authorization middleware** — protected routes require a valid JWT; only event owners can update/delete their own events
- **SQLite database** — lightweight, file-based storage with foreign key relationships between users, events, and registrations

## Tech Stack

- **Language:** Go
- **Web framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** SQLite (via `modernc.org/sqlite`)
- **Auth:** JSON Web Tokens ([golang-jwt/jwt](https://github.com/golang-jwt/jwt))
- **Password hashing:** bcrypt (`golang.org/x/crypto/bcrypt`)

## Project Structure

```
├── db/            # Database connection and table setup
├── middlewares/   # Authorization middleware (JWT verification)
├── models/        # Data models and database queries (User, Event)
├── routes/        # Route handlers and route registration
├── utils/         # Helper functions (password hashing, JWT generation/validation)
├── api-test/      # Sample .http request files for manual testing
└── main.go        # Application entry point
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.20 or later installed

### Setup

1. Clone the repository
   ```bash
   git clone https://github.com/adityasingh688-maker/go-event-registration-api.git
   cd go-event-registration-api
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Run the server
   ```bash
   go run main.go
   ```

The server will start at `http://localhost:8080`, and a SQLite database file (`api.db`) will be created automatically on first run.

> **Note:** This project currently uses a hardcoded JWT secret for simplicity. Before deploying anywhere beyond local testing, move it to an environment variable.

## API Endpoints

### Public

| Method | Endpoint         | Description         |
|--------|------------------|----------------------|
| POST   | `/signup`        | Create a new user account |
| POST   | `/login`         | Log in and receive a JWT |
| GET    | `/events`        | Get all events |
| GET    | `/events/:id`    | Get a single event by ID |

### Authenticated (requires `Authorization` header with a valid JWT)

| Method | Endpoint                     | Description                          |
|--------|-------------------------------|---------------------------------------|
| POST   | `/events`                    | Create a new event |
| PUT    | `/events/:id`                | Update an event (owner only) |
| DELETE | `/events/:id`                | Delete an event (owner only) |
| POST   | `/events/:id/register`       | Register for an event |
| DELETE | `/events/:id/register`       | Cancel registration for an event |

### Example: Sign up

```http
POST http://localhost:8080/signup
Content-Type: application/json

{
    "email": "test@example.com",
    "password": "yourpassword"
}
```

### Example: Create an event

```http
POST http://localhost:8080/events
Content-Type: application/json
Authorization: <your-jwt-token>

{
    "name": "Test Event",
    "description": "An example event",
    "location": "A test location",
    "dateTime": "2025-01-01T15:30:00.000Z"
}
```

More example requests are available in the `api-test/` folder as `.http` files, usable with the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for VS Code.

## Future Improvements

- Move the JWT secret to an environment variable
- Add input validation (e.g. email format checks)
- Add unit tests for models and utils
- Add pagination for the events list

## License

This project is for learning and portfolio purposes.