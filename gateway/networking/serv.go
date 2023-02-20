package networking

import (
	"fmt"
	"log"
	"net"

	"github.com/Vivena/Toy-QKD-IKE/gateway/core"
)

type Serv struct {
	port   string
	saList map[string]core.IkeSA
}

func (i *Serv) NewServ(port string) *Serv {
	return &Serv{port: port}
}

func (i *Serv) Start() error {
	maxUDPSize := 65507

	buffer := make([]byte, maxUDPSize)
	conn, err := net.ListenUDP("udp", i.port)
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
