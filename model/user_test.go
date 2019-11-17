package model

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUser_HashPassword(t *testing.T) {
	user := &User{Password: "password"}
	if err := user.HashPassword(); err != nil {
		t.Error("Error while hashing")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password")); err != nil {
		t.Error("Password did not matched hash")
	}

}

func TestUser_ComparePassword(t *testing.T) {
	user := &User{Password: "password"}
	user.HashPassword()
	if err := user.ComparePassword("password"); err != nil {
		t.Error("comparepassword returned error")
	}
}
