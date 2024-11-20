package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	room := models.Room{
		Name:        "Test Room",
		Description: "This is a test room",
		Status:      "available",
	}
	jsonData, _ := json.Marshal(room)
	req, _ := http.NewRequest("POST", "/rooms", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Test Room", data["name"])

	config.Setup()
}

func TestGetRooms(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/rooms", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])

	config.Setup()
}

func TestCreateReservation(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	reservation := models.Reservation{
		RoomID:   1,
		Date:     "2024-11-22",
		Time:     "03:00 PM",
		PlayerID: 1,
	}
	jsonData, _ := json.Marshal(reservation)
	req, _ := http.NewRequest("POST", "/reservations", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(1), data["room_id"])
	assert.Equal(t, float64(1), data["player_id"])

	config.Setup()
}

func TestGetReservations(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/reservations?room_id=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])

	config.Setup()
}
