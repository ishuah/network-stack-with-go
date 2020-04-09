package main

import "encoding/binary"

type UDPDatagram []byte

func (d UDPDatagram) SourcePort() uint16 {
	return binary.BigEndian.Uint16(d[0:2])
}

func (d UDPDatagram) DestinationPort() uint16 {
	return binary.BigEndian.Uint16(d[2:4])
}

func (d UDPDatagram) Length() uint16 {
	return binary.BigEndian.Uint16(d[4:6])
}

func (d UDPDatagram) Body() string {
	return string(d[8:d.Length()])
}
