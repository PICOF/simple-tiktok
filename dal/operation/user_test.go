package operation

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	ret, err := GetUser(context.Background(), 1)
	if err != nil {
		return
	}
	fmt.Println(ret)
}

func TestGetUserByName(t *testing.T) {
	ret, err := GetUserByName(context.Background(), "Jiyeon")
	if err != nil {
		return
	}
	fmt.Println(ret)
}

func TestWriteUserInfo(t *testing.T) {
	info := &TUserInfo{
		Username: "Jiyeon",
		Password: "1234567",
	}
	ret, err := WriteUserInfo(context.Background(), info)
	if err != nil {
		return
	}
	fmt.Println(ret)
}
func TestGetWorkCount(t *testing.T) {
	count, err := GetWorkCount(context.Background(), 24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
