package payloads

import (
	"unsafe"

	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
)

type QKD_KeyID_payload struct {
	Header headers.QKD_KeyID_header
	key_ID string
}

func New_QKD_KeyID_payload(key_ID string, output *QKD_KeyID_payload) (uint16, error) {

	output.key_ID = key_ID

	output.Header = *headers.New_QKD_KeyID_header()

	return uint16(len(key_ID)) + uint16(unsafe.Sizeof(output.Header)), nil
}

func (p *QKD_KeyID_payload) Key_ID() string {
	return p.key_ID
}
