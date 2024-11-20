# Game Room Management System API Documentation

## General Information

- Base URL: `http://localhost:8080`
- All requests and responses are in JSON format.

## Endpoints

### 1. `/rooms`

#### `GET /rooms`
- **Description**: Retrieve all game room information.
- **Example Request**:
  ```bash
  curl -X GET http://localhost:8080/rooms
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "name": "Room A",
        "description": "This is room A",
        "status": "available"
      },
      {
        "id": 2,
        "name": "Room B",
        "description": "This is room B",
        "status": "booked"
      }
    ]
  }
  ```

#### `POST /rooms`
- **Description**: Create a new game room.
- **Request Body**:
  ```json
  {
    "name": "Room D",
    "description": "This is room D",
    "status": "available"
  }
  ```
- **Example Request**:
  ```bash
  curl -X POST http://localhost:8080/rooms -H "Content-Type: application/json" -d '{"name": "Room D", "description": "This is room D", "status": "available"}'
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 4,
      "name": "Room D",
      "description": "This is room D",
      "status": "available"
    }
  }
  ```

### 2. `/rooms/{id}`

#### `GET /rooms/{id}`
- **Description**: Retrieve detailed information about a specific room.
- **Example Request**:
  ```bash
  curl -X GET http://localhost:8080/rooms/1
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 1,
      "name": "Room A",
      "description": "This is room A",
      "status": "available"
    }
  }
  ```

#### `PUT /rooms/{id}`
- **Description**: Update information about a specific room.
- **Request Body**:
  ```json
  {
    "name": "Room A Updated",
    "description": "Updated description for Room A",
    "status": "booked"
  }
  ```
- **Example Request**:
  ```bash
  curl -X PUT http://localhost:8080/rooms/1 -H "Content-Type: application/json" -d '{"name": "Room A Updated", "description": "Updated description for Room A", "status": "booked"}'
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 1,
      "name": "Room A Updated",
      "description": "Updated description for Room A",
      "status": "booked"
    }
  }
  ```

#### `DELETE /rooms/{id}`
- **Description**: Delete a specific room.
- **Example Request**:
  ```bash
  curl -X DELETE http://localhost:8080/rooms/1
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": "Room deleted successfully"
  }
  ```

### 3. `/reservations`

#### `GET /reservations`
- **Description**: Query game room reservations. Supports optional query parameters to filter results.
- **Query Parameters**:
  - `room_id` (optional): The room ID to filter reservations.
  - `date` (optional): The date to filter reservations.
  - `limit` (optional): The maximum number of reservations to return.
- **Example Request**:
  ```bash
  curl -X GET "http://localhost:8080/reservations?room_id=1&date=2024-11-18&limit=2"
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "room_id": 1,
        "date": "2024-11-18",
        "time": "10:00 AM",
        "player_id": 1
      }
    ]
  }
  ```

#### `POST /reservations`
- **Description**: Create a new reservation for a game room.
- **Request Body**:
  ```json
  {
    "room_id": 1,
    "date": "2024-11-22",
    "time": "03:00 PM",
    "player_id": 3
  }
  ```
- **Example Request**:
  ```bash
  curl -X POST http://localhost:8080/reservations -H "Content-Type: application/json" -d '{"room_id": 1, "date": "2024-11-22", "time": "03:00 PM", "player": "David"}'
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 4,
      "room_id": 1,
      "date": "2024-11-22",
      "time": "03:00 PM",
      "player_id": 3
    }
  }
  ```
```