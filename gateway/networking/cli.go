package networking

import (
	"context"
	"math/rand"
	"net"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
	"github.com/Vivena/Toy-QKD-IKE/gateway/payload"
)

type Cli struct {
	addr   *net.UDPAddr
	qkd    *crypto.QKD
	saList map[uint32]core.IkeSA
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
	c.qkd = crypto.NewQKD(qkd_ip, qkd_port, "test")
	return &c, nil
}

func (c *Cli) get_SPI() uint32 {
	// Todo: See if we add a max number of retry
	for {
		res := rand.Uint32()
		if _, ok := c.saList[res]; !ok {
			return res
		}
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
	overall_size := 0

	var key_Payload payload.QKD_KeyID_payload
	var tf_payload payload.QKD_Transform_payload

	tmp, err := payload.New_QKD_KeyID_payload(key.Key_id, &key_Payload)
	if err != nil {
		goto END
	}
	//TODO:
	//key_Payload.header.Set_next_payload()
	key_Payload.header.Set_payload_len(tmp)
	key_Payload.header.Set_version()
	key_Payload.header.Set_mode_N()
	key_Payload.header.Set_device_id(c.qkd.saeID)
	key_Payload.header.Set_key_id_len(len(key_Payload.key_ID))

	overall_size += tmp

END:
	// TODO: Zero the key
	return err
}
