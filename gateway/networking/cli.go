package networking

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"unsafe"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
	"github.com/Vivena/Toy-QKD-IKE/gateway/payloads"
)

type Cli struct {
	addr     *net.UDPAddr
	qkd      *crypto.QKD
	listLock sync.Mutex
	SaList   map[uint32]core.IkeSA
}

func (c *Cli) QKD() *crypto.QKD {
	return c.qkd
}

func NewCli(sa_ip string, qkd_ip string, qkd_port string) (*Cli, error) {
	// We set the addr for the other end-point
	var c Cli

	ServerAddr, err := net.ResolveUDPAddr("udp", sa_ip+":"+constants.SA_port)
	if err != nil {
		return nil, err
	}
	c.addr = ServerAddr

	// We setup the qkd
	c.qkd = crypto.NewQKD(qkd_ip, qkd_port, 1)
	c.SaList = make(map[uint32]core.IkeSA)
	return &c, nil
}

func (c *Cli) get_SPI() uint32 {
	// Todo: See if we add a max number of retry
	ike_sa := core.IkeSA{Name: 0, State: "IKE_SA_INIT"}
	for {
		res := rand.Uint32()
		c.listLock.Lock()
		_, ok := c.SaList[res]
		if !ok {
			c.SaList[res] = ike_sa
			c.listLock.Unlock()
			return res
		}
		c.listLock.Unlock()
	}
}

func (c *Cli) Init_IKE_SA() error {
	ctx := context.Background()
	// We first get the key from the QKD
	key, err := c.qkd.GetKey(ctx, 256)

	// TODO: error handeling
	if err != nil {
		panic(err)
	}
	overall_size := uint32(0)

	var key_Payload payloads.QKD_KeyID_payload

	tmp, err := payloads.New_QKD_KeyID_payload(key.Key_id, &key_Payload)
	if err != nil {
		return err
	}

	key_Payload.Header.Set_next_payload(constants.QKD_KEY_ID)
	key_Payload.Header.Set_payload_len(tmp)
	key_Payload.Header.Set_version()
	key_Payload.Header.Set_mode_N()
	key_Payload.Header.Set_device_id(c.QKD().SaeID)
	key_Payload.Header.Set_key_id_len(uint16(len(key_Payload.Key_ID())))

	overall_size += uint32(tmp)

	tf_payload := payloads.New_Transform_payload()
	tf_payload.Set_is_last()

	overall_size += uint32(unsafe.Sizeof(tf_payload))

	ike_header := headers.New_IKE_Header()

	ike_header.SetDefaultFlag()
	ike_header.SetIsInitFlag()
	ike_header.Set_INIT_SPI(c.get_SPI())
	ike_header.Set_RESP_SPI(0)
	ike_header.Set_next_payload(constants.QKD_PAYLOAD)
	ike_header.Set_message_id(constants.IKE_SA_INIT)

	overall_size += uint32(unsafe.Sizeof(ike_header))
	ike_header.Set_length(overall_size)

	// TODO: Zero the key
	return err
}
