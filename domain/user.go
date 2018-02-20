package domain

type User struct {
	Id   uint `gorm:"primary_key"`
	Name string
	Age  int
}
