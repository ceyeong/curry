package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User :
type User struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name" bson:"name"`
	Email             string             `json:"email" bson:"email"`
	Password          string             `json:"-" bson:"password"`
	CreatedAt         time.Time          `json:"createAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt         time.Time          `json:"updateAt,omitempty" bson:"updatedAt,omitempty"`
	PasswordUpdatedAt time.Time          `json:"passwordUpdatedAt,omitempty" bson:"passwordUpdatedAt,omitempty"`
}

// HashPassword : hash user password
func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println(string(hash))
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

// ComparePassword : compares the password. returns nil if match
func (user *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// NewUser returns new User instance.
func NewUser() *User {
	return new(User)
}
