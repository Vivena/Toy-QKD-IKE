package networking

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
	"github.com/Vivena/Toy-QKD-IKE/gateway/headers"
)

type Serv struct {
	Conn     net.PacketConn
	listLock sync.Mutex
	saList   map[uint32]core.IkeSA
}

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
	maxUDPSize := 65507

	defer i.Conn.Close()
	buffer := make([]byte, maxUDPSize)
	for {
		n, remoteAddr, err := i.Conn.ReadFrom(buffer[0:])
		if err != nil {
			log.Fatalf("Error:%s\n", err)
		}
		fmt.Println(buffer[0:n])
		if len(buffer) >= constants.IKE_HEADER_SIZE {
			go serve(i, remoteAddr, buffer)
		}
	}
}

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

// func handle_IKE_SA_INIT(sa *core.saIkeSA) {

// }

//TODO: Later, we currently do not handle the case where the server init the rekey
// func handle_IKE_SA_INIT_RESP()  {

// }

//
// func IKE_AUTH_RESP()  {

// }
func IKE_AUTH() {

}
