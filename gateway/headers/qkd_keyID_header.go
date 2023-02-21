package headers

import (
	"encoding/binary"
	"fmt"
	"io"
)

// QKD_KeyID_header This struct holds information about the QKD Key ID header
type QKD_KeyID_header struct {
	Next_payload  uint8
	Reserved      uint8
	Payload_len   uint16
	Version       uint8
	Flags         uint8
	QKD_device_id uint16
	KeyID_len     uint16
}

func New_QKD_KeyID_header() *QKD_KeyID_header {
	var res QKD_KeyID_header
	res.set_critical_bit()
	return &res
}

// To avoid a man-in-the-middle attack downgrading the negotiated security level, the Critical bit
// must be set to 1.

// Set_critical_bit: set the critical bit to 1
func (h *QKD_KeyID_header) set_critical_bit() {
	h.Reserved |= 1 << 7
}

// Verify_critical_bit: verify the critical bit
func (h *QKD_KeyID_header) Verify_critical_bit() bool {
	return (h.Reserved & (1 << 7)) != 0
}

// Get_mode_F: Get the fallback field
func (h *QKD_KeyID_header) Get_mode_F() bool {
	return (h.Flags & (1 << 7)) != 0
}

// Set_mode_F: Set the fallback field to fallback mode
func (h *QKD_KeyID_header) Set_mode_F() {
	h.Flags |= (1 << 7)
}

// Set_mode_N: Set the fallback field to normal mode
func (h *QKD_KeyID_header) Set_mode_N() {
	h.Flags &= ^uint8(0x1 << 7)
}

func (h *QKD_KeyID_header) Set_Next_payload(Next_payload uint8) {
	h.Next_payload = Next_payload
}

func (h *QKD_KeyID_header) Set_Payload_len(Payload_len uint16) {
	h.Payload_len = Payload_len
}

func (h *QKD_KeyID_header) Set_Version() {
	h.Version = 1
}

func (h *QKD_KeyID_header) Set_device_id(QKD_device_id uint16) {
	h.QKD_device_id = QKD_device_id
}

func (h *QKD_KeyID_header) Set_key_id_len(KeyID_len uint16) {
	h.KeyID_len = KeyID_len
}

// TODO: add error handeling

func (h *QKD_KeyID_header) Read_header(buf io.Reader) {
	err := binary.Read(buf, binary.LittleEndian, h)
	if err != nil {
		fmt.Println("failed to Read:", err)
		return
	}
}

func (h *QKD_KeyID_header) Write_header(buf io.Writer) {
	err := binary.Write(buf, binary.LittleEndian, h)
	if err != nil {
		fmt.Println("failed to Read:", err)
		return
	}
}
