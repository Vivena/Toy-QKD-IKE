package payloads

// We do not handle TF attributes
type QKD_Transform_payload struct {
	last          uint8
	reserved1     uint8
	transformLen  uint16
	transformType uint8
	reserved2     uint8
	transformID   uint16
}

func New_Transform_payload() *QKD_Transform_payload {
	return &QKD_Transform_payload{reserved1: 0, transformLen: 8, transformType: 241, reserved2: 0, transformID: 1}
}

func (t *QKD_Transform_payload) Set_is_last() {
	t.last = 0
}
func (t *QKD_Transform_payload) Set_is_not_last() {
	t.last = 3
}
