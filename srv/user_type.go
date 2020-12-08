package main

import (
	"crypto/subtle"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	// Admin role
	Admin UserRole = iota
	// Researcher role
	Researcher
	// Subject role
	Subject
)

// UserRole type binding to integer
type UserRole uint8

// User for this app
type User struct {
	gorm.Model `json:"-"`
	ID         uint64    `json:"id"  gorm:"primaryKey"`
	Username   string    `json:"username"  gorm:"unique"`
	Password   string    `json:"password,omitempty"`
	Birthday   time.Time `json:"birthday" gorm:"index"`
	Role       UserRole  `json:"role"`
}

// ErrorMessage returned after smth got wrong
type ErrorMessage struct {
	Code    uint        `json:"code"`
	Payload interface{} `json:"payload"`
	Message string      `json:"message"`
}

// GetHashed user with hashed password
func (u *User) GetHashed() (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return &User{ID: u.ID, Username: u.Username, Password: string(bytes), Role: u.Role}, err
}

// Check the password
func (u *User) Check(v *User) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(v.Password)) == nil
}

// SetAdminUser updates the admin on start
func SetAdminUser(db *gorm.DB, c Config) {
	// Create Admin User
	admin := &User{ID: 1, Username: "admin", Password: c.AdminPassword, Role: Admin}
	adminHashed, err := admin.GetHashed()
	if err != nil {
		panic(err)
	}
	if err := db.Save(adminHashed).Error; err != nil {
		panic(err)
	}
}

// DatabaseBasicAuth returns a basic auth function with db curry
func DatabaseBasicAuth(db *gorm.DB) func(username, password string, c echo.Context) (bool, error) {
	return func(username, password string, c echo.Context) (bool, error) {
		var dbUser User
		rqUser := &User{Username: username, Password: password}
		if err := db.Table("users").Where("username = ?", username).Scan(&dbUser).Error; err != nil {
			return false, err
		}
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(rqUser.Username), []byte(dbUser.Username)) == 1 &&
			dbUser.Check(rqUser) {
			return true, nil
		}
		return false, nil
	}
}
