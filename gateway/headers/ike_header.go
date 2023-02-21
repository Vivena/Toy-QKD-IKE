package headers

import (
	"bytes"
	"encoding/binary"
)

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
	Flags           uint8
	MessageID       uint32
	Length          uint32
}

func New_IKE_Header() *IKE_Header {
	return &IKE_Header{Version: uint8(0x22)}
}

func (h *IKE_Header) flagSetReserved() {
	h.Flags &= 0x38
}

func (h *IKE_Header) flagSetVersion() {
	h.Flags &= 0xEF
}

func (h *IKE_Header) SetDefaultFlag() {
	h.flagSetReserved()
	h.flagSetVersion()
}

func (h *IKE_Header) SetIsInitFlag() {
	h.Flags |= 1 << 3
}

func (h *IKE_Header) IsInit() bool {
	return !(h.Flags&(0x1<<3) == 0)
}

func (h *IKE_Header) SetIsRespFlag() {
	h.Flags |= 1 << 5
}

func (h *IKE_Header) IsResp() bool {
	return !(h.Flags&(0x1<<5) == 0)
}

func (h *IKE_Header) Set_INIT_SPI(spi uint32) {
	h.IKE_SA_INIT_SPI = spi
}

func (h *IKE_Header) Set_RESP_SPI(spi uint32) {
	h.IKE_SA_RESP_SPI = spi
}

func (h *IKE_Header) Set_next_payload(payload uint8) {
	h.Next_payload = payload
}

func (h *IKE_Header) Set_exchange_type(exchange_type uint8) {
	h.Exchange_type = exchange_type
}

func (h *IKE_Header) Set_message_id(message_id uint32) {
	h.MessageID = message_id
}

func (h *IKE_Header) Set_length(length uint32) {
	h.Length = length
}

func IKE_Header_Parse(header_bytes []byte) (*IKE_Header, error) {
	var h IKE_Header
	reader := bytes.NewReader(header_bytes)

	err := binary.Read(reader, binary.LittleEndian, &h.IKE_SA_INIT_SPI)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.IKE_SA_RESP_SPI)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Next_payload)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Version)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Exchange_type)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Flags)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.MessageID)
	if err != nil {
		goto ERR
	}
	err = binary.Read(reader, binary.LittleEndian, &h.Length)
	if err != nil {
		goto ERR
	}

	return &h, nil
ERR:
	return nil, err
}
