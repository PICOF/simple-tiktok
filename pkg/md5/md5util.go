package md5

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GetFileMd5(content []byte) (string, error) {
	md5h := md5.New()
	_, err := io.Copy(md5h, bytes.NewReader(content))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
