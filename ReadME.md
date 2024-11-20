# RESTful API Project by Gin

## Overview

This project is a RESTful API developed using the Gin framework in Go. It covers various functionalities followed by Interview2024, including player management, level management, game room management, reservations, game log collection, endless challenges, and payment processing(unfinished).

## Project Structure

```plaintext
project-root/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── config/
│   ├── config.go          // Configuration file, such as data file paths and database settings
├── controllers/
│   ├── player_controller.go // Handles player management logic
│   ├── level_controller.go  // Handles level management logic
│   ├── room_controller.go   // Handles game room management logic
│   ├── reservation_controller.go // Handles reservation management logic
│   ├── log_controller.go    // Handles game log management logic
│   ├── challenge_controller.go // Handles endless challenge system logic
├── models/
│   ├── player.go           // Player model, defines structure
│   ├── level.go            // Level model, defines structure
│   ├── room.go             // Room model, defines structure
│   ├── reservation.go      // Reservation model, defines structure
│   ├── log.go              // Log model, defines structure
│   ├── challenge.go        // Challenge model, defines structure
│   ├── payment.go          // Payment model, defines structure
├── repositories/
│   ├── player_repository.go // Abstract interface and implementation for player data
│   ├── level_repository.go  // Abstract interface and implementation for level data
│   ├── room_repository.go   // Abstract interface and implementation for room data
│   ├── reservation_repository.go // Abstract interface and implementation for reservation data
│   ├── log_repository.go    // Abstract interface and implementation for log data
│   ├── challenge_repository.go // Abstract interface and implementation for challenge data
│   ├── common_repository.go // Common repository functions shared across other repositories
├── routes/
│   ├── routes.go          // Defines all routes
├── services/
│   ├── player_service.go   // Player service, includes business logic
│   ├── level_service.go    // Level service, includes business logic
│   ├── room_service.go     // Room service, includes business logic
│   ├── reservation_service.go // Reservation service, includes business logic
│   ├── log_service.go      // Log service, includes business logic
│   ├── challenge_service.go // Challenge service, includes business logic
├── tests/
│   ├── player_test.go     // Player Management System API test script
│   ├── room_test.go       // Game Room Management System API test script
│   ├── setup_router.go    // Common router functions shared across other tests script
├── utils/
│   ├── response.go        // Handles common response formatting
│   ├── file_utils.go      // Handles reading and writing to files
├── docs/
│   ├── Player Management System API Documentation.md // Detailed description of all API endpoints
│   ├── Game Room Management System API Documentation.md // Detailed description of all API endpoints
│   ├── Endless Challenge System API Documentation.md // Detailed description of all API endpoints
│   ├── Payment Processing System API Documentation.md // Detailed description of all API endpoints
├── data/
    ├── players.json         // Sample player data for testing
    ├── rooms.json           // Sample room data for testing
    ├── reservations.json    // Sample reservation data for testing
    ├── logs.json            // Sample log data for testing
    ├── challenges.json      // Sample challenge data for testing
```

## Features

### Player Management
- Add, update, retrieve, and delete player information.
- Endpoints: `/players`, `/players/{id}`.

### Level Management
- Manage game levels.
- Endpoints: `/levels`, `/levels/{id}`.

### Room Management
- Create, update, delete, and get information about game rooms.
- Endpoints: `/rooms`, `/rooms/{id}`.

### Reservations
- Reserve game rooms and manage reservations.
- Endpoints: `/reservations`.

### Game Log Collection
- Collect and query player activity logs such as login, logout, and challenges.
- Endpoints: `/logs`.

### Endless Challenge System
- Players can participate in endless challenges, and the challenge results are updated periodically.
- Endpoints: `/challenges`, `/challenges/results`.

### Payment Processing System(unfinished)
- Process payments with various methods such as credit card, bank transfer, third-party, and blockchain.
- Endpoints: `/payments`, `/payments/{id}`.

## How to Run the Project

### Prerequisites
- Go 1.23.3
- Docker
- Docker Compose

### Steps to Run
1. Clone the repository.
   ```bash
   git clone https://github.com/smbb1234/interview_Steve_Nov19th.git
   cd interview_Steve_Nov19th
   ```
2. Build Docker image.
   ```bash
   docker build -t gin-api .
   ```
3. Run the services using Docker Compose.
   ```bash
   docker-compose up
   ```
4. The application will be accessible at `http://localhost:8080`.

### Running Tests
- To run the API tests:
  ```bash
  go test ./tests/
  ```

## API Documentation
Detailed API documentation is available in the `docs/` folder:
- **[Player Management System](/docs/Player%20Management%20SystemAPI%20Documentation.md)**: Describes player-related endpoints.
- **[Game Room Management System](/docs/Game%20Room%20Management%20System%20API%20Documentation.md)**: Describes room and reservation management.
- **[Endless Challenge System](/docs/Endless%20Challenge%20System%20API%20Documentation.md)**: Provides details about participating in challenges.
- **[Game Log Collector](/docs/Game%20Log%20Collector%20API%20Documentation.md)**: Recording every action of the player.
- **Payment Processing System(unfinished)**: Describes payment endpoints and different payment methods.

## Configuration
- **config/config.go**: Contains application configuration details, including database settings and data file paths.

## Data Storage
- The application can store data using only JSON files(MySQL support is unfinished).
- By default, JSON files are used.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Author
- **Steve**: Main developer of the project.