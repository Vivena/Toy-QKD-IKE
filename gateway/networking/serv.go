package networking

import (
	"fmt"
	"log"
	"net"

	"github.com/Vivena/Toy-QKD-IKE/gateway/constants"
	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
)

type Serv struct {
	addr   *net.UDPAddr
	saList map[string]core.IkeSA
}

func NewServ() (*Serv, error) {
	var s Serv
	addr, err := net.ResolveUDPAddr("udp", constants.SA_port)
	if err != nil {
		return nil, err
	}
	s.addr = addr
	return &s, nil
}

func (i *Serv) Start() error {
	maxUDPSize := 65507

	buffer := make([]byte, maxUDPSize)
	conn, err := net.ListenUDP("udp", i.addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		_, remoteAddr, err := conn.ReadFromUDP(buffer[0:])
		if err != nil {
			log.Fatalf("Error:%s\n", err)
		}

		go serve(conn, remoteAddr, buffer)
	}
}

func serve(conn net.PacketConn, remoteAddr net.Addr, buffer []byte) {
	responseStr := fmt.Sprintf("toto")
	conn.WriteTo([]byte(responseStr), remoteAddr)
}
