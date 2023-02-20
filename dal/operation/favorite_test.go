package operation

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestFavorite(t *testing.T) {
	FavoriteAction(context.Background(), 24, 5, true)
	FavoriteAction(context.Background(), 24, 5, false)
}
func TestGetFavoriteCount(t *testing.T) {
	count, err := GetFavoriteCount(context.Background(), 24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
