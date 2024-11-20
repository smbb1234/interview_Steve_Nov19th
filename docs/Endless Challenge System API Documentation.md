# Endless Challenge System API Documentation

## General Information

- Base URL: `http://localhost:8080`
- All requests and responses are in JSON format.

## Endpoints

### 1. `/challenges`

#### `POST /challenges`
- **Description**: Players can participate in a challenge by providing their player ID and participation count. Each participation attempt requires a fixed participation count and has a cooldown of 1 minute.
- **Request Body**:
  ```json
  {
    "player_id": 1,
    "amount": 20.01
  }
  ```
- **Example Request**:
  ```bash
  curl -X POST http://localhost:8080/challenges -H "Content-Type: application/json" -d '{"player_id": 1, "amount": 20.01}'
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": {
      "id": 11,
      "player_id": 1,
      "participation_count": 1,
      "status": "Win",
      "timestamp": "2024-11-18 10:00:00"
    }
  }
  ```
- **Error Response** (If player participates again within the cooldown time):
  ```json
  {
    "success": false,
    "error": "You can only participate in a challenge once per minute"
  }
  ```

### 2. `/challenges/results`

#### `GET /challenges/results`
- **Description**: Retrieve the most recent challenge results, including details such as challenge ID, player ID, participation count, status, and timestamp.
- **Query Parameters**: None.
- **Example Request**:
  ```bash
  curl -X GET "http://localhost:8080/challenges/results"
  ```
- **Example Response**:
  ```json
  {
    "success": true,
    "data": [
      {
        "id": 1,
        "player_id": 2,
        "participation_count": 3,
        "status": "Win",
        "timestamp": "2024-11-17 15:30:00"
      },
      {
        "id": 2,
        "player_id": 1,
        "participation_count": 1,
        "status": "Lose",
        "timestamp": "2024-11-17 14:00:00"
      }
    ]
  }
  ```

## Logic

1. **Participation Requirements**
   - Players must provide their player ID and participation count.
   - Each participation has a fixed amount (`participation_count`) and must be greater than 0.
   - Players are only allowed to participate once every minute.

2. **Challenge Logic**
   - Each challenge lasts for 30 seconds.
   - Players can join challenges at most once per minute.
   - After 30 seconds, the player has a 1% chance per participation count to win the entire prize pool.
   - Players can continue to participate, with higher participation increasing the probability of winning.

## Status Values

- The `status` field for challenges can have one of the following values:
  - `Win`: The player won the challenge.
  - `Lose`: The player lost the challenge.
  - `Pending`: The challenge outcome is still pending.

These statuses represent the results of the player's challenge attempts.

## Notes

- Ensure that the `player_id` provided in requests is valid and corresponds to an existing player.
- The `timestamp` is generated automatically when a new challenge is created.
- The cooldown logic ensures that players cannot participate in challenges more frequently than once per minute.

