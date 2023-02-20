//go:build integration
// +build integration

package crypto

import (
	"context"
	"testing"
)

func GetKey_test(t *testing.T) {

	ctx := context.Background()

	qkd := NewQKD("127.0.0.1", "8000", "test")
	print("test")
	key, err := qkd.GetKey(ctx, 256)

	if err != nil {
		println(err)
	}
	print(key.Key_id)
}
