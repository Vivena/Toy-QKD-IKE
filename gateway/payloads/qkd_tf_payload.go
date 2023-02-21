package payloads

// We do not handle TF attributes for the moment

type QKD_Transform_payload struct {
	last          uint8
	reserved1     uint8
	transformLen  uint16
	transformType uint8
	reserved2     uint8
	transformID   uint16
}

// New_Transform_payload: create a basic QKD_Transform_payload
func New_Transform_payload() *QKD_Transform_payload {
	return &QKD_Transform_payload{reserved1: 0, transformLen: 8, transformType: 241, reserved2: 0, transformID: 1}
}

//TODO: create getters for all attributes

// Set_is_last: set last attribute to 0 to signify this is the last Transform_payload
func (t *QKD_Transform_payload) Set_is_last() {
	t.last = 0
}

// Set_is_last: set last attribute to 3 to signify this is not the last Transform_payload
func (t *QKD_Transform_payload) Set_is_not_last() {
	t.last = 3
}

// TODO: write parser for QKD_Transform_payload struct
func Parse_QKD_Transform_payload(payload []byte) (*QKD_Transform_payload, error) {
	return nil, nil
}
