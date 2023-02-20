package crypto

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"
)

func TestNewQKD(t *testing.T) {
	want := "127.0.0.1"
	qkd := NewQKD(want, "8000", "test")
	if qkd.url != want {
		t.Errorf("got %q, wanted %q", qkd.url, want)
	}
}

func TestGetKey(t *testing.T) {

	ctx := context.Background()

	qkd := NewQKD("127.0.0.1", "8000", "test")

	key, err := qkd.GetKey(ctx, 256)

	if err != nil {
		t.Errorf("%w", err)
	}

	if !(key.Key != nil) {
		t.Errorf("key is empty")
	}

	if key.Key_id == "" {
		t.Errorf("keyID is empty")
	}

	if len(key.Key) != 32 {
		t.Errorf("Incorrect key len %d, expected %d", len(key.Key), 32)
	}

}

func TestGetKeyWithID(t *testing.T) {

	ctx := context.Background()

	qkd := NewQKD("127.0.0.1", "8000", "test")

	key1, err := qkd.GetKey(ctx, 256)

	if err != nil {
		t.Errorf("%w", err)
	}

	key2, err := qkd.GetKeyWithID(ctx, key1.Key_id)

	if err != nil {
		t.Errorf("%w", err)
	}

	if !(key2.Key != nil) {
		t.Errorf("key is empty")
	}

	if key2.Key_id == "" {
		t.Errorf("keyID is empty")
	}

	key2.Key, err = hex.DecodeString(key2.Key_tmp)

	if err != nil {
		t.Errorf("%w", err)
	}

	if len(key2.Key) != 32 {
		t.Errorf("Incorrect key len %d, expected %d", len(key2.Key), 32)
	}

	if key1.Key_id != key2.Key_id {
		t.Errorf("Incorrect keyID %s, expected %s", key2.Key_id, key1.Key_id)
	}

	key1.Key, err = hex.DecodeString(key1.Key_tmp)

	if err != nil {
		t.Errorf("%w", err)
	}

	if !bytes.Equal(key1.Key, key2.Key) {
		t.Errorf("Incorrect key %x, expected %x", key1.Key, key2.Key)
	}

}
