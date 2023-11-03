package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	Email    string `gorm:"size: 255; not null;unique" json:"email"`
	Username string `gorm:"size: 255; not null;" json:"username"`
	Password string `gorm:"size: 255; not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
