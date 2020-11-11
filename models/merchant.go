package models

import (
	"time"
)

type (
	Merchant struct {
		ID        int       `json:"id"`
		Name      string    `name:"name"`
		Telp      string    `json:"telp"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
