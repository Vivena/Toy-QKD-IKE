package networking

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
	"unsafe"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
	"github.com/Vivena/Toy-QKD-IKE/gateway/payloads"
)

// Cli: basic information to be able to run a IKE instance
type Cli struct {
	// addr: UDP address of the server
	addr *net.UDPAddr
	// qkd:	information to connect to the assigned qkd
	qkd *crypto.QKD
	// listLock: allows to lock the saList
	listLock sync.Mutex
	// SaList: List of all the IKE SA registered in the IKE instance
	SaList map[uint32]core.IkeSA
}

// QKD: getter for the qkd field
func (c *Cli) QKD() *crypto.QKD {
	return c.qkd
}

// TODO: set getter for all the Cli fields

// TODO: We might want to decouple the NewCli from udp address solving to allow for the instance to start
// without a server to connect to.
// NewCli: create a new Cli instance and resolve the upd address for the server we want to connect to
func NewCli(sa_ip string, qkd_ip string, qkd_port string) (*Cli, error) {
	// We set the addr for the other end-point
	var c Cli

	ServerAddr, err := net.ResolveUDPAddr("udp", sa_ip+":"+constants.SA_port)
	if err != nil {
		return nil, err
	}
	c.addr = ServerAddr

	// We setup the qkd based on the info provided
	// TODO: once we don't run a dummy qkd,we will need to ask the user to provide
	// the qkd id instead of using 1 as a default qkd id
	c.qkd = crypto.NewQKD(qkd_ip, qkd_port, 1)
	c.SaList = make(map[uint32]core.IkeSA)
	return &c, nil
}

// get_SPI: get a new SPI to identify the new SA connection
func (c *Cli) get_SPI() uint32 {
	// TODO: see if we add a max number of retry
	// TODO: see if we also create a list of known SPI instead of only looking at
	// SPI used for initializer point of view (Do we create a central DB later to handle that?)
	ike_sa := core.IkeSA{Name: 0, State: "IKE_SA_INIT"}
	for {
		res := rand.Uint32()
		c.listLock.Lock()
		defer c.listLock.Unlock()
		_, ok := c.SaList[res]
		if !ok {
			c.SaList[res] = ike_sa
			return res
		}
	}
}

// create_Packet_content_Init_IKE_SA: create the Init_IKE_SA message
func (c *Cli) create_Packet_content_Init_IKE_SA() ([]byte, error) {
	// TODO: only set timeout if asked to
	// ctx is the context for the connection with the QKD
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*constants.Timeout))
	defer cancel()

	// We first get the key from the QKD (see if we want to allow for different key size later)
	key, err := c.qkd.GetKey(ctx, 256)
	// TODO: error handeling
	if err != nil {
		panic(err)
	}

	// Starting now we will track the overall size of the message as we built it
	overall_size := uint32(0)

	var key_Payload payloads.QKD_KeyID_payload
	fmt.Println("Get QKD key")
	tmp, err := payloads.New_QKD_KeyID_payload(key.Key_id, &key_Payload)
	if err != nil {
		return nil, err
	}

	fmt.Println("Create Key Payload")
	key_Payload.Header.Set_Next_payload(constants.QKD_KEY_ID)
	key_Payload.Header.Set_Payload_len(tmp)
	key_Payload.Header.Set_Version()
	key_Payload.Header.Set_mode_N()
	key_Payload.Header.Set_device_id(c.QKD().SaeID)
	key_Payload.Header.Set_key_id_len(uint16(len(key_Payload.Key_ID())))

	overall_size += uint32(tmp)

	fmt.Println("Create Transform Payload")
	tf_payload := payloads.New_Transform_payload()
	tf_payload.Set_is_last()

	overall_size += uint32(unsafe.Sizeof(tf_payload))

	fmt.Println("Create IKE Header")
	ike_header := headers.New_IKE_Header()

	ike_header.SetDefaultFlag()
	ike_header.SetIsInitFlag()
	ike_header.Set_INIT_SPI(c.get_SPI())
	ike_header.Set_RESP_SPI(0)
	ike_header.Set_next_payload(constants.QKD_PAYLOAD)
	ike_header.Set_message_id(0)

	overall_size += uint32(unsafe.Sizeof(ike_header))
	ike_header.Set_length(overall_size)

	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, ike_header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, tf_payload)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, key_Payload.Header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, []byte(key_Payload.Key_ID()))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Cli) Init_IKE_SA_Reply(conn net.Conn) error {
	return nil
}

func (c *Cli) Init_IKE_SA() error {

	content, err := c.create_Packet_content_Init_IKE_SA()
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, c.addr)
	if err != nil {
		fmt.Println("Listen failed:", err.Error())
		return err
	}

	//close the connection
	defer conn.Close()
	fmt.Println("Sending the content")
	fmt.Println(content)
	_, err = conn.Write(content)
	if err != nil {
		fmt.Println("Write data failed:", err.Error())
		return err
	}
	fmt.Println("Done.")
	// TODO: Zero the key
	return nil
}
