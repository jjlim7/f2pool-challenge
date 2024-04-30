package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DNSLog struct {
	gorm.Model
	ID        int32          `gorm:"primaryKey;autoIncrement:true"`
	Domain    string         `gorm:"not null"`
	IPv4IPs   pq.StringArray `gorm:"type:varchar(255)[];not null"` // Can have multiple IP Addr?
	CreatedAt time.Time      `gorm:"autoCreateTime:true"`
}
