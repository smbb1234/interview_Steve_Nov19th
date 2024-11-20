package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"steveInterviewMod/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayers(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/players", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])

	config.Setup()
}

func TestCreatePlayer(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	player := map[string]interface{}{
		"name":     "Test Player",
		"level_id": 1,
		"others":   "Extra information about Test Player",
	}
	jsonData, _ := json.Marshal(player)
	req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Test Player", data["name"])
	assert.Equal(t, float64(1), data["level_id"])

	config.Setup()
}

func TestGetPlayerByID(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/players/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Alice", data["name"])
	assert.Equal(t, float64(1), data["level_id"])

	config.Setup()
}

func TestUpdatePlayer(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	player := map[string]interface{}{
		"name":     "Updated Player",
		"level_id": 2,
		"others":   "Updated information about Player",
	}
	jsonData, _ := json.Marshal(player)
	req, _ := http.NewRequest("PUT", "/players/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Updated Player", data["name"])
	assert.Equal(t, float64(2), data["level_id"])

	config.Setup()
}

func TestDeletePlayer(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/players/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Player deleted successfully", response["data"])

	config.Setup()
}

func TestGetLevels(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/levels", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])

	config.Setup()
}

func TestCreateLevel(t *testing.T) {
	SetupTestPath()

	router := setupRouter()
	level := map[string]interface{}{
		"name": "Advanced",
	}
	jsonData, _ := json.Marshal(level)
	req, _ := http.NewRequest("POST", "/levels", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Advanced", data["name"])

	config.Setup()
}
