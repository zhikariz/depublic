package entity

type Transaction struct {
	ID              int64
	UserID          int64
	PaymentID       int64
	TransactionDate string
	User            *User
	Payment         *TransactionPayment
	Details         []TransactionDetail
}

func (Transaction) TableName() string {
	return "transactions"
}
