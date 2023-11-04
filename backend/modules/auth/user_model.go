package auth

import (
	"html"
	"strings"

	"github.com/Enhancifire/Petrolium/backend/modules/storage"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	UserID   string `gorm:"size: 255; not null; unique; primary_key" json:"user_id"`
	Email    string `gorm:"size: 255; not null; unique;" json:"email"`
	Username string `gorm:"size: 255; not null;" json:"username"`
	Password string `gorm:"size: 255; not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	err := storage.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func ReturnModels() []any {
	return []any{
		User{},
	}
	// storage.Migrate(User{})
}
