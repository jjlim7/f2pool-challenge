package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"backend/db"
	"backend/models"
)

var conn, _ = db.ConnectDatabase()

// RootHandler handles requests to the root endpoint "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	info := models.AppInfo{
		Version:    os.Getenv("APP_VERSION"),
		Date:       time.Now().Unix(),
		Kubernetes: isRunningInKubernetes(),
	}
	sendJSONResponse(w, http.StatusOK, info)
}

// send JSON response with the specified status code
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		log.Printf("[Error] %v", err)
	}
}

// check if the application is running in a Kubernetes environment
func isRunningInKubernetes() bool {
	_, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token")
	return err == nil
}
