Here is your content formatted as Markdown:

```markdown
# Walkit

A fitness workout tracking API built with Go, Gin framework, and PocketBase. This project is based on the [Fitness Workout Tracker](https://roadmap.sh/projects/fitness-workout-tracker) project from roadmap.sh.

## Features

- User Authentication (JWT)
- Exercise Management
- Workout Planning
- User Profile Management
- API Documentation with Swagger
- PocketBase Database Integration

## Prerequisites

- Go 1.19 or higher
- PocketBase
- Air (for hot reload)
- Swag (for API documentation)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rowjay007/walkit.git
   cd walkit
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Install development tools:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   go install github.com/cosmtrek/air@latest
   ```

4. Set up environment variables (create a `.env` file):
   ```env
   PORT=8080
   POCKETBASE_URL=http://127.0.0.1:8090
   JWT_SECRET=your_jwt_secret
   ```

5. Start PocketBase:
   ```bash
   ./pocketbase serve
   ```

## Development

Start the development server with hot reload:

```bash
make run
```

Other available commands:

```bash
make build    # Build the application
make test     # Run tests
make docs     # Generate Swagger documentation
make clean    # Clean build artifacts
```

## API Documentation

Access the Swagger UI at:

```plaintext
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/forgot-password` - Request password reset
- `POST /api/v1/auth/reset-password` - Reset password

### Users
- `GET /api/v1/users/me` - Get current user profile
- `PATCH /api/v1/users/me` - Update current user
- `DELETE /api/v1/users/me` - Delete current user
- `GET /api/v1/users` - List users
- `GET /api/v1/users/:id` - Get user by ID
- `PATCH /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Exercises
- `POST /api/v1/exercises` - Create exercise
- `GET /api/v1/exercises` - List exercises
- `GET /api/v1/exercises/:id` - Get exercise by ID
- `PATCH /api/v1/exercises/:id` - Update exercise
- `DELETE /api/v1/exercises/:id` - Delete exercise

### Workouts
- `POST /api/v1/workouts` - Create workout
- `GET /api/v1/workouts/:id` - Get workout by ID
- `PATCH /api/v1/workouts/:id` - Update workout
- `DELETE /api/v1/workouts/:id` - Delete workout

## Project Structure

```plaintext
walkit/
├── cmd/
│   └── server/
│       └── main.go
├── config/
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── routes/
│   └── service/
├── pkg/
│   ├── logger/
│   └── util/
├── docs/
├── .air.toml
├── .env
├── Makefile
└── README.md
```

## Database

This project uses PocketBase as its database solution. PocketBase is an open-source backend consisting of embedded database with realtime subscriptions, built-in auth, and file storage.

### PocketBase Admin UI

Access the PocketBase Admin UI at:

```plaintext
http://127.0.0.1:8090/_/
```

### Collections

The following collections are used in the project:

- users
- exercises
- workouts
- workout_schedules

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
