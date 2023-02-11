package bcrypt

import (
	"context"
	"testing"
)

func TestBcrypt(t *testing.T) {
	pwd := "1919810"
	password, err := EncryptPassword(context.Background(), pwd)
	if err != nil {
		return
	}
	println(string(password))
	println(len(string(password)))
	password2, err := EncryptPassword(context.Background(), pwd)
	println(string(password2))
	println(ComparePassword(string(password), pwd))
}
