package main

import (
	"net"
	"log"
	dns "github.com/davecheney/dns"
)

func main() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 5353,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	err = socket.JoinGroup(net.IPv4(224, 0, 0, 251))
	if err != nil {
		log.Fatal(err)

	}
	for {
		buff := make([]byte, 4096)
		read, err := socket.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		msg := &dns.Msg{}
		msg.Unpack(buff[:read])
		log.Println(msg.String())
	}
}
