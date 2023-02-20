package headers

import (
	"encoding/binary"
	"fmt"
	"io"
)

// QKD_KeyID_header This struct holds information about the QKD Key ID header
type QKD_KeyID_header struct {
	next_payload  uint8
	reserved      uint8
	payload_len   uint16
	version       uint8
	flags         uint8
	qkd_device_id uint16
	keyID_len     uint16
}

// To avoid a man-in-the-middle attack downgrading the negotiated security level, the Critical bit
// must be set to 1.

// Set_critical_bit: set the critical bit to 1
func (h *QKD_KeyID_header) Set_critical_bit() {
	h.reserved = 1 << 7
}

// Verify_critical_bit: verify the critical bit
func (h *QKD_KeyID_header) Verify_critical_bit() bool {
	return (h.reserved & (1 << 7)) != 0
}

// Get_mode_F: Get the fallback field
func (h *QKD_KeyID_header) Get_mode_F() bool {
	return (h.flags & (1 << 7)) != 0
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
