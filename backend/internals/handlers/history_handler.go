// File: internal/handlers/history_handler.go
package handlers

import (
	"net/http"

	"backend/models"
)

// HistoryHandler handles requests to the /v1/history endpoint
func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the latest 20 saved queries from the database
	var logs []models.Query
	conn.Order("created_at desc").Limit(20).Find(&logs)

	sendJSONResponse(w, http.StatusOK, logs)
}
