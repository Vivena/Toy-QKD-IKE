package payloads

import (
	"errors"
	"unsafe"

	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
)

// We do not handle TF attributes
type QKD_Transform_payload struct {
	last          uint8
	reserved1     uint8
	transformLen  uint16
	transformType uint8
	reserved2     uint8
	transformID   uint16
}

func New_QKD_KeyID_payload(key_ID []byte, output *QKD_KeyID_payload) (uint16, error) {

	output.key_ID = make([]byte, len(key_ID))
	n := copy(output.key_ID, key_ID)
	if n != len(key_ID) {
		return 0, errors.New("error copying key ")
	}
	output.header = *headers.New_QKD_KeyID_header()

	return uint16(n) + uint16(unsafe.Sizeof(output.header)), nil
}
