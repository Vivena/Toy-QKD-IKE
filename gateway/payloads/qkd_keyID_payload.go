package payloads

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
)

type QKD_KeyID_payload struct {
	Header        headers.QKD_KeyID_header
	Key_ID_String string
}

// New_QKD_KeyID_payload: create a basic QKD_KeyID_payload using a basic QKD_KeyID_header
func New_QKD_KeyID_payload(key_ID string, output *QKD_KeyID_payload) (uint16, error) {

	output.Key_ID_String = key_ID

	output.Header = *headers.New_QKD_KeyID_header()

	return uint16(len(key_ID)) + uint16(unsafe.Sizeof(output.Header)), nil
}

// Key_ID: getter for Key_ID
func (p *QKD_KeyID_payload) Key_ID() string {
	return p.Key_ID_String
}

// We parse the QKD_KeyID_payload in two time:
//		- first we parse the fixed size payload header which gives us
//		  the Key_ID_String length
//		- then we parse the Key_ID_String

// Parse_QKD_KeyID_payload_header: parce the header of QKD_KeyID payload
func Parse_QKD_KeyID_payload_header(payload []byte) (*QKD_KeyID_payload, error) {

	var h QKD_KeyID_payload
	reader := bytes.NewReader(payload)

	err := binary.Read(reader, binary.LittleEndian, &h.Header.Next_payload)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.Reserved)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.Payload_len)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.Version)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.Flags)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.QKD_device_id)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Header.KeyID_len)
	if err != nil {
		return nil, err
	}

	return &h, nil
}

// Parse_QKD_KeyID_payload_Key_name: parse Key_ID_String part of the QKD_KeyID payload
func Parse_QKD_KeyID_payload_Key_name(payload []byte, payload_size uint16, h *QKD_KeyID_payload) error {
	reader := bytes.NewReader(payload)
	tmp := make([]byte, payload_size)
	err := binary.Read(reader, binary.LittleEndian, &tmp)
	if err != nil {
		return err
	}

	h.Key_ID_String = string(tmp)
	return nil
}
