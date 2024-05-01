// File: internal/handlers/lookup_handler.go
package handlers

import (
	"backend/models"
	"log"
	"net"
	"net/http"
)

// LookupHandler handles requests to the /v1/tools/lookup endpoint
func LookupHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Domain parameter is required", http.StatusBadRequest)
		return
	}

	// Perform DNS lookup for IPv4 addresses
	ips, err := LookupIPv4(domain)
	if err != nil {
		http.Error(w, "Error performing DNS lookup", http.StatusInternalServerError)
		return
	}

	// Log the successful query to the database
	query, err := logQuery(domain, ips, r.RemoteAddr)
	if err != nil {
		http.Error(w, "Error logging query to database", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, http.StatusOK, query)
}

// lookupIPv4 performs a DNS lookup and returns IPv4 addresses for a given domain
func LookupIPv4(domain string) ([]string, error) {
	addrs, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, addr := range addrs {
		if ip4 := addr.To4(); ip4 != nil {
			addresses = append(addresses, ip4.String())
		}
	}
	log.Printf("lookupIPv4 %v", addresses)
	return addresses, nil
}

// logQuery logs the successful DNS lookup query and its result to the database
func logQuery(domain string, ips []string, clientip string) (models.Query, error) {

	// Create a model.Query object
	query := models.Query{
		Addresses: ips,
		ClientIP:  clientip,
		Domain:    domain,
	}

	// Insert query and result into the database
	err := conn.Create(&query).Error
	if err != nil {
		log.Printf("Error inserting Query %v", err)
		return models.Query{}, err
	}

	return query, nil
}
