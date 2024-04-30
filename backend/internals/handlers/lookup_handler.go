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
	ips, err := lookupIPv4(domain)
	if err != nil {
		http.Error(w, "Error performing DNS lookup", http.StatusInternalServerError)
		return
	}

	// Log the successful query to the database
	err = logQuery(domain, ips)
	if err != nil {
		http.Error(w, "Error logging query to database", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"domain":  domain,
		"ipv4ips": ips,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

// lookupIPv4 performs a DNS lookup and returns IPv4 addresses for a given domain
func lookupIPv4(domain string) ([]string, error) {
	addrs, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ipv4ips []string
	for _, addr := range addrs {
		if ip4 := addr.To4(); ip4 != nil {
			ipv4ips = append(ipv4ips, ip4.String())
		}
	}
	log.Printf("lookupIPv4 %v", ipv4ips)
	return ipv4ips, nil
}

// logQuery logs the successful DNS lookup query and its result to the database
func logQuery(domain string, ipv4ips []string) error {

	dnsLog := models.DNSLog{
		Domain:  domain,
		IPv4IPs: ipv4ips,
	}

	// Insert query and result into the database
	err := conn.Create(&dnsLog).Error
	if err != nil {
		log.Printf("Error inserting DNSLog %v", err)
		return err
	}

	return nil
}
