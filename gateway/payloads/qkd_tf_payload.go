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
