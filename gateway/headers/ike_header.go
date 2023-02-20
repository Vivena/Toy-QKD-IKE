package headers

const (
	IKE_SA_INIT = 34
	IKE_AUTH    = 35
)

type IKE_Header struct {
	IKE_SA_INIT_SPI uint32
	IKE_SA_RESP_SPI uint32
	Next_payload    uint8
	Version         uint8
	Exchange_type   uint8
	flags           uint8
	MessageID       uint32
	length          uint32
}

func (h *IKE_Header) flagSetReserved() {
	h.flags &= 0x38
}

func (h *IKE_Header) flagSetVersion() {
	h.flags &= 0xEF
}

func (h *IKE_Header) setDefaultFlag() {
	h.flagSetReserved()
	h.flagSetVersion()
}

// TODO
func NewIKEHeader() *IKE_Header {
	return nil
}

func (h *IKE_Header) SetIsInitFlag() {
	h.flags |= 1 << 3
}

func (h *IKE_Header) IsInit() bool {
	return !(h.flags&(0x1<<3) == 0)
}

func (h *IKE_Header) SetIsRespFlag() {
	h.flags |= 1 << 5
}

func (h *IKE_Header) IsResp() bool {
	return !(h.flags&(0x1<<5) == 0)
}
