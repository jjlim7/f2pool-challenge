package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// AppInfo represents information about the application.
type AppInfo struct {
	Version    string `json:"version"`
	Date       int64  `json:"date"`
	Kubernetes bool   `json:"kubernetes"`
}

type Query struct {
	gorm.Model
	ID        int32          `gorm:"primaryKey;autoIncrement:true"`
	ClientIP  string         `json:"client_ip"`
	Domain    string         `json:"domain"`
	Addresses pq.StringArray `gorm:"type:varchar(255)[];not null" json:"addresses"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true" json:"created_at"`
}
