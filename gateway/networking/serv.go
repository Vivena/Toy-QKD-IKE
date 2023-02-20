package networking

import (
	"fmt"
	"log"
	"net"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
)

type Serv struct {
	Conn   net.PacketConn
	saList map[string]core.IkeSA
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
		go serve(i.Conn, remoteAddr, buffer)
	}
}

func serve(conn net.PacketConn, remoteAddr net.Addr, buffer []byte) {

	// conn.WriteTo([]byte(responseStr), remoteAddr)
}
