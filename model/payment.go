package model

import "time"

type Payment struct {
	PaymentId         string    `json:"payment_id" db:"payment_id"`
	SenderId          string    `json:"sender_id" db:"sender_id"`
	ReceiverId        string    `json:"receiver_id" db:"receiver_id"`
	Amount            string    `json:"amount" db:"amount"`
	BankAccountNumber string    `json:"bank_account_number" db:"bank_account_number"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}
