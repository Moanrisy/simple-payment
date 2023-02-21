package model

import "time"

type Customer struct {
	CustomerId string    `json:"customer_id" db:"customer_id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	Balance    string    `json:"balance" db:"balance"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
