package model

import "time"

type Merchant struct {
	MerchantId string    `json:"merchant_id" db:"merchant_id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	Balance    string    `json:"balance" db:"balance"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
