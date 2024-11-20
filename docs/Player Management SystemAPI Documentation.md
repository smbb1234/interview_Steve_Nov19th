# Player Management System API Documentation

## General Information

- Base URL: `http://localhost:8080`
- All requests and responses are in JSON format.

## Endpoints

### 1. `/players`

#### `GET /players`
- Description: Retrieve all player information.
- Example Request:
  ```bash
  curl -X GET http://localhost:8080/players
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "name": "Alice",
        "level_id": 1,
        "other": "Extra information about Alice"
      },
      {
        "id": 2,
        "name": "Bob",
        "level_id": 2,
        "other": "Additional details about Bob"
      }
    ]
  }
  ```

#### `POST /players`
- Description: Register a new player.
- Request Body:
  ```json
  {
    "name": "Charlie",
    "level_id": 1,
    "other": "Extra information about Charlie"
  }
  ```
- Example Request:
  ```bash
  curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"name": "Charlie", "level": {"id": 1, "name": "Beginner"}, "other": "Extra information about Charlie"}'
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": {
      "id": 3,
      "name": "Charlie",
      "level_id": 1,
      "other": "Extra information about Charlie"
    }
  }
  ```

### 2. `/players/{id}`

#### `GET /players/{id}`
- Description: Retrieve detailed information about a specific player.
- Example Request:
  ```bash
  curl -X GET http://localhost:8080/players/1
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": {
      "id": 1,
      "name": "Alice",
      "level_id": 1,
      "other": "Extra information about Alice"
    }
  }
  ```

#### `PUT /players/{id}`
- Description: Update information about a specific player.
- Request Body:
  ```json
  {
    "name": "Alice Updated",
    "level_id": 2,
    "other": "Updated information about Alice"
  }
  ```
- Example Request:
  ```bash
  curl -X PUT http://localhost:8080/players/1 -H "Content-Type: application/json" -d '{"name": "Alice Updated", "level": {"id": 2, "name": "Intermediate"}, "other": "Updated information about Alice"}'
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": {
      "id": 1,
      "name": "Alice Updated",
      "level_id": 2,
      "other": "Updated information about Alice"
    }
  }
  ```

#### `DELETE /players/{id}`
- Description: Delete a specific player.
- Example Request:
  ```bash
  curl -X DELETE http://localhost:8080/players/1
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": "Player deleted successfully"
  }
  ```

### 3. `/levels`

#### `GET /levels`
- Description: Retrieve all level information.
- Example Request:
  ```bash
  curl -X GET http://localhost:8080/levels
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "name": "Beginner"
      },
      {
        "id": 2,
        "name": "Intermediate"
      }
    ]
  }
  ```

#### `POST /levels`
- Description: Add a new level.
- Request Body:
  ```json
  {
    "name": "Advanced"
  }
  ```
- Example Request:
  ```bash
  curl -X POST http://localhost:8080/levels -H "Content-Type: application/json" -d '{"name": "Advanced"}'
  ```
- Example Response:
  ```json
  {
    "success": true,
    "data": {
      "id": 3,
      "name": "Advanced"
    }
  }
  ```
```