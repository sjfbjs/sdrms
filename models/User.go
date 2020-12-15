package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID      int
	Name    string
	Age     int
	AddTime int
}

func (User) TableName() string {
	return "user"
}
