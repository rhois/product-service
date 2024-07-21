package entity

import (
	"time"
)

type Suppliers struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
	Products    []Products
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
