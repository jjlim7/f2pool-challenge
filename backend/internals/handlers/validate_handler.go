package handlers

import (
	"encoding/json"
	"net"
	"net/http"
)

type ValidateIPRequest struct {
	IP string `json:"ip"`
}

type ValidateIPResponse struct {
	Status bool `json:"status"`
}

// ValidateHandler handles requests to the /v1/tools/validate endpoint
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var req ValidateIPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate IP address format
	if req.IP == "" {
		http.Error(w, "IP parameter is required", http.StatusBadRequest)
		return
	}

	resp := ValidateIPResponse{
		Status: IsValidIPv4(req.IP),
	}
	sendJSONResponse(w, http.StatusOK, resp)
}

// check if given string is a valid IPv4 address
func IsValidIPv4(ip string) bool {
	return net.ParseIP(ip) != nil && net.ParseIP(ip).To4() != nil
}
