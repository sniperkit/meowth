package models

import (
	"github.com/weeq/meowth/lib"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(30);unique_index" json:"username"'`
	Email    string `gorm:"type:varchar(50);unique_index" json:"email"'`
	Password string `gorm:"type:varchar(88);unique_index" json:"-"'`
	*lib.Model
}