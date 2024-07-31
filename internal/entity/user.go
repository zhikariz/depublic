package entity

type User struct {
	ID       int64
	Username string
	Password string
	Name     string
}

func (User) TableName() string {
	return "users"
}
