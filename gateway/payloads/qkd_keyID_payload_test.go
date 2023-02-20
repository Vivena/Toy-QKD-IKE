package payloads

import (
	"context"
	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
	"testing"
)

func Test_New_QKD_KeyID_payload(t *testing.T) {
	var payload QKD_KeyID_payload
	ctx := context.Background()
	qkd := crypto.NewQKD("127.0.0.1", "8000", "test")

	key, err := qkd.GetKey(ctx, 256)
	if err != nil {
		t.Errorf(" %w", err)
	}
	n, err := New_QKD_KeyID_payload(key.Key, &payload)
	if err != nil {
		t.Errorf(" %w", err)
	}

	if n != 42 {
		t.Errorf("payload size is %d, expected 42", n)
	}
}
