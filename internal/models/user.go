package models

type User struct {
	ID   int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
}

func (u *User) TableName() string {
	return "users"
}
