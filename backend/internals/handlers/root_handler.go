package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"backend/db"
	"backend/models"
)

var conn, _ = db.ConnectDatabase()

// RootHandler handles requests to the root endpoint "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	info := models.AppInfo{
		Version:    "0.1.0",
		Date:       time.Now().Unix(),
		Kubernetes: isRunningInKubernetes(),
	}
	sendJSONResponse(w, http.StatusOK, info)
}

// sendJSONResponse sends a JSON response with the specified status code
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// isRunningInKubernetes checks if the application is running in a Kubernetes environment
func isRunningInKubernetes() bool {
	// Implement your logic here to detect Kubernetes environment
	return false
}
