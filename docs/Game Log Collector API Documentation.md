# Game Log Collector API Documentation

## General Information

- Base URL: `http://localhost:8080`
- All requests and responses are in JSON format.

## Endpoints

### 1. `/logs`

#### `GET /logs`
- **Description**: Retrieve game logs, with optional filters for player ID, action type, time range, and limit.
- **Query Parameters**:
  - `player_id` (optional): The player ID to filter logs.
  - `action` (optional): The action type to filter logs. Allowed values:
    - `register`
    - `login`
    - `logout`
    - `enter_room`
    - `leave_room`
    - `join_challenge`
    - `challenge_result`
  - `start_time` (optional): Start of the time range for the logs (in `YYYY-MM-DD HH:MM:SS` format).
  - `end_time` (optional): End of the time range for the logs (in `YYYY-MM-DD HH:MM:SS` format).
  - `limit` (optional): The maximum number of logs to return.
- **Example Request**:
  ```bash
  curl -X GET "http://localhost:8080/logs?player_id=1&action=login&start_time=2024-11-10%2010:00:00&end_time=2024-11-20%2010:00:00&limit=5"
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "player_id": 1,
        "action": "login",
        "timestamp": "2024-11-15 15:30:00",
        "details": "Details for action 1"
      },
      {
        "id": 2,
        "player_id": 1,
        "action": "logout",
        "timestamp": "2024-11-16 16:00:00",
        "details": "Details for action 2"
      }
    ]
  }
  ```

#### `POST /logs`
- **Description**: Create a new game action log.
- **Request Body**:
  ```json
  {
    "player_id": 1,
    "action": "login",
    "details": "Player logged in successfully."
  }
  ```
- **Example Request**:
  ```bash
  curl -X POST http://localhost:8080/logs -H "Content-Type: application/json" -d '{"player_id": 1, "action": "login", "details": "Player logged in successfully."}'
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 11,
      "player_id": 1,
      "action": "login",
      "timestamp": "2024-11-18 10:00:00",
      "details": "Player logged in successfully."
    }
  }
  ```

## Action Types

- The `action` parameter can have one of the following values:
  - `register`: Register
  - `login`: Login
  - `logout`: Logout
  - `enter_room`: Enter Room
  - `leave_room`: Leave Room
  - `join_challenge`: Join Challenge
  - `challenge_result`: Challenge Result

These actions represent different types of player activities that can be logged for analysis and tracking purposes.

## Notes

- Ensure that the `player_id` provided in requests is valid and corresponds to an existing player.
- The `timestamp` is generated automatically when a new log is created.
- The time range for filtering logs (`start_time` and `end_time`) should be in `YYYY-MM-DD HH:MM:SS` format to avoid errors.