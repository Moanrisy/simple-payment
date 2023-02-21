package model

type Bank struct {
	BankId            string `json:"bank_id" db:"bank_id"`
	BankAccountNumber string `json:"bank_account_number" db:"bank_account_number"`
	Balance           string `json:"balance" db:"balance"`
}
