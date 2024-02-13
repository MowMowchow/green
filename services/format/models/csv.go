package models

import "time"

type RawCSV struct {
	AccountType string `json:"account_type,omitempty"`;
	AccountNumber string `json:"account_number,omitempty"`;
	Name string `json:"name,omitempty"`;
	TransactionDate time.Time `json:"transaction_date,omitempty"`;
	Description1 string `json:"description1,omitempty"`;
	Description2 string `json:"description2,omitempty"`;
	Description3 string `json:"description3,omitempty"`;
	Cad float32 `json:"cad,omitempty"`;
}