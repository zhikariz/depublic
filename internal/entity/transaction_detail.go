package entity

type TransactionDetail struct {
	ID            int64
	TransactionID int64
	ProductID     int64
	Qty           int64
	Price         float64
	Product       *Product
}
