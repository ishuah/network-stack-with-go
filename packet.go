package main

import "net"

type Packet []byte

func (p Packet) Version() byte {
	return p[0] >> 4
}

func (p Packet) IHL() byte {
	return p[0] & 0xF
}

func (p Packet) Protocol() byte {
	return p[9]
}

func (p Packet) SourceIP() net.IP {
	return net.IP(p[12:16])
}

func (p Packet) DestinationIP() net.IP {
	return net.IP(p[16:20])
}

func (p Packet) Options() []byte {
	return p[20:]
}
