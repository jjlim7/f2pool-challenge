package models

// AppInfo represents information about the application.
type AppInfo struct {
	Version    string `json:"version"`
	Date       int64  `json:"date"`
	Kubernetes bool   `json:"kubernetes"`
}
