package bcrypt

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(ctx context.Context, password string) (result []byte, err error) {
	result, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while encrypting the password: %v", err)
		return nil, err
	}
	return
}
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
