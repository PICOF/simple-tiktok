package operation

import (
	"context"
	"testing"
)

func TestFavorite(t *testing.T) {
	FavoriteAction(context.Background(), 24, 5, true)
	FavoriteAction(context.Background(), 24, 5, false)
}
