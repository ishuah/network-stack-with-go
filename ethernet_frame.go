package main

import "net"

type EthernetFrame []byte

func (ef EthernetFrame) Destination() net.HardwareAddr {
	return net.HardwareAddr(ef[:6:6])
}

func (ef EthernetFrame) Source() net.HardwareAddr {
	return net.HardwareAddr(ef[6:12:12])
}

func (ef EthernetFrame) Data() []byte {
	return ef[14 : len(ef)-4]
}
