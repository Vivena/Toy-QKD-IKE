package payloads

// TODO: change the timemout to no timeout when required in crypto package
// Possibled value for the QKD_fallback_payload fallback field
const (
	WAIT_QKD      uint16 = 1
	DIFFIEHELLMAN uint16 = 2
	CONTINUE      uint16 = 3
)

type QKD_fallback_payload struct {
	next_payload uint8
	reserved     uint8
	payload_len  uint16
	version      uint8
	flags        uint8
	fallback     uint16
}

// TODO: create New function for QKD_fallback_payload to init a basic QKD_fallback_payload

// TODO: create getters

// TODO: create QKD_fallback_payload parser
