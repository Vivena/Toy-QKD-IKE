package networking

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
)

// Serv: basic information to be able to run a IKE instance
type Serv struct {
	// Conn: on-going connection for the instance
	Conn net.PacketConn
	// qkd:
	qkd *crypto.QKD
	// listLock: allows to lock the saList
	listLock sync.Mutex
	// saList: List of all the IKE SA registered in the IKE instance
	saList map[uint32]core.IkeSA
}

// QKD: getter for qkd field
func (c *Serv) QKD() *crypto.QKD {
	return c.qkd
}

// TODO: set all the getters for Serv

// NewServ: create a new Serv instance and init the server connection
func NewServ() (*Serv, error) {
	var s Serv
	conn, err := net.ListenPacket("udp", ":"+constants.SA_port)
	if err != nil {
		return nil, err
	}
	s.Conn = conn
	return &s, nil
}

func (i *Serv) Start() error {
	// TODO: do proper multiple read of the buffer instead of using a giant buffer per connection
	// We want to do the read in two times:
	//		- get the IKE header (fixed size)
	//		- use the IKE header to try to create a buffer of the correct size, and read the rest of the message

	// maxUDPSize is set to the maximum size for udp.
	maxUDPSize := 65507

	// Don't forget to close the connection when we are done
	defer i.Conn.Close()
	// As it is required for each messages to fit inside a single udp paquet, we always get the full message
	buffer := make([]byte, maxUDPSize)

	// Basic udp server looping to get messages
	for {
		n, remoteAddr, err := i.Conn.ReadFrom(buffer[0:])
		if err != nil {
			log.Fatalf("Error:%s\n", err)
		}
		fmt.Println(buffer[0:n])

		// at minimum the message need to countain an IKE header, if that's not the case,
		// the message is obviously malformed, so we drop it.
		if len(buffer) >= constants.IKE_HEADER_SIZE {
			// We will handle the message in its own go routine to be able to continue receiving
			// new messages while we handle this one.
			go serve(i, remoteAddr, buffer)
		}
	}
}

// serve: function that handle a message (this is where the logic for message handeling is)
func serve(i *Serv, remoteAddr net.Addr, buffer []byte) {
	ike_header, err := headers.IKE_Header_Parse(buffer[0:constants.IKE_HEADER_SIZE])
	if err != nil {
		return
	}

	//TODO: verify payload length match the rest of the buffer
	i.listLock.Lock()
	sa, ok := i.saList[ike_header.IKE_SA_INIT_SPI]
	i.listLock.Unlock()
	if ok {
		// This SA already exist
		//We need to take a lock for it
		sa.SALock.Lock()
		if sa.State == "IKE_SA_INIT" {
			if ike_header.Exchange_type != constants.IKE_SA_INIT || ike_header.IsInit() {
				return
			}
			//TODO handle the IKE_SA_INIT packet

		} else if sa.State == "IKE_SA_INIT_RESP" {
			//TODO: for the moment, the server is only passif
			// if ike_header.Exchange_type != constants.IKE_AUTH || !ike_header.IsInit() {
			// 	return
			// }
			//TODO handle the IKE_SA_INIT
		} else if sa.State == "IKE_AUTH" {
			//TODO: we need to verify we have a key before allowing fallback

		} else if sa.State == "QKD_KEY_UP" {
			//In that case, we already have a key, and we are expecting a rekey
			if ike_header.Exchange_type != constants.IKE_SA_INIT || !ike_header.IsInit() {
				return
			}
		}
		defer sa.SALock.Unlock()

	} else {
		//SA does not exist
		//if exchange type is KE_SA_INIT we create it
		if ike_header.Exchange_type == constants.IKE_SA_INIT {
			res := rand.Uint32()
			sa := core.IkeSA{Name: res, State: "IKE_SA_INIT_RESP"}
			i.listLock.Lock()
			i.listLock.Unlock()

			i.saList[ike_header.IKE_SA_INIT_SPI] = sa
			i.listLock.Unlock()

		}
	}
}

// TODO: finish writing the functions for all the possible cases
// IKE_SA_INIT - rekey or not
// IKE_SA_INIT_RESP - rekey or not
// IKE_AUTH - rekey or not
// IKE_AUTH_RESP - rekey or not

// func handle_IKE_SA_INIT(sa *core.saIkeSA) {
//
// }

//TODO: Later, we currently do not handle the case where the server init the rekey
// func handle_IKE_SA_INIT_RESP()  {
//
// }

// func IKE_AUTH() {
//
// }

// func IKE_AUTH_RESP()  {
//
// }
