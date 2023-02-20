package payloads

const (
	WAIT_QKD 		uint16 = 1
	DIFFIE-HELLMAN	uint16 = 2
	CONTINUE		uint16 = 3 
)

type QKD_fallback_payload struct {
	next_payload uint8
	reserved     uint8
	payload_len  uint16
	version      uint8
	flags        uint8
	fallback     uint16
}
