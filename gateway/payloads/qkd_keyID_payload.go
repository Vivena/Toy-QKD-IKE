package payloads

import (
	"github.com/Vivena/gateway/Toy-QKD-IKE/headers"
)

type QKD_KeyID_payload struct {
	header headers.QKD_KeyID_header
	key_ID []byte
}
