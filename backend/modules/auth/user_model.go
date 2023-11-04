package auth

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/Enhancifire/Petrolium/backend/modules/storage"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ReturnModels() []any {
	return []any{
		User{},
	}
}

type User struct {
	UserID    string    `gorm:"size: 255; not null; unique; primary_key" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `gorm:"size: 255; not null; unique;" json:"email"`
	Username  string    `gorm:"size: 255; not null; unique" json:"username"`
	Password  string    `gorm:"size: 255; not null;" json:"password"`
}

func GetUserByID(id string) (*User, error) {
	var user User
	err := storage.DB.Model(&User{}).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	user.Password = ""
	return &user, nil
}

func (u *User) SaveUser() (*User, error) {
	u.UserID = uuid.New().String()
	err := storage.DB.Model(&User{}).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println(string(hashedPass))
	u.Password = string(hashedPass)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func getUsers() ([]User, error) {
	var users []User
	err := storage.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email, password string) (string, error) {
	var user User

	if err := storage.DB.Model(&User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return "", err
	}

	if err := VerifyPassword(password, user.Password); err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateToken(user.UserID)

	if err != nil {
		return "", err
	}

	return token, nil
}
