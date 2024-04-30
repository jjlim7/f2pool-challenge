package handlers

import (
	"net/http"
)

// HealthHandler handles requests to the health endpoint "/health"
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, "OK")
}
