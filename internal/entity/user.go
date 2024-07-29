package entity

type User struct {
	ID   int64
	Name string
}

func (User) TableName() string {
	return "users"
}
