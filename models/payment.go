package models

type PaymentMethod string

const (
	CreditCard   PaymentMethod = "CreditCard"
	BankTransfer PaymentMethod = "BankTransfer"
	ThirdParty   PaymentMethod = "ThirdParty"
	Blockchain   PaymentMethod = "Blockchain"
)

type Payment struct {
	ID            uint          `json:"id"`
	Method        PaymentMethod `json:"method"`
	Amount        float64       `json:"amount"`
	Details       string        `json:"details"`
	Status        string        `json:"status"`
	TransactionID string        `json:"transaction_id"`
	Timestamp     string        `json:"timestamp"`
}
