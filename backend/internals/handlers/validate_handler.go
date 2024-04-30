package handlers

import (
	"net"
	"net/http"
)

// ValidateHandler handles requests to the /v1/tools/validate endpoint
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "IP parameter is required", http.StatusBadRequest)
		return
	}

	resp := map[string]interface{}{
		"ip":      ip,
		"isValid": isValidIPv4(ip),
	}
	sendJSONResponse(w, http.StatusOK, resp)
}

// check if given string is a valid IPv4 address
func isValidIPv4(ip string) bool {
	return net.ParseIP(ip) != nil && net.ParseIP(ip).To4() != nil
}
