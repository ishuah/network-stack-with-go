package main

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"unsafe"
)

func htons(h int) (n int) {
	a := uint16(42)
	if *(*byte)(unsafe.Pointer(&a)) == 42 {
		a = uint16(h)
		n = int(a>>8 | a<<8)
	} else {
		n = h
	}
	return
}

func main() {
	fd, err := unix.Socket(unix.AF_PACKET, unix.SOCK_RAW, htons(unix.ETH_P_ALL))
	if err != nil {
		log.Fatal(err)
	}

	defer unix.Close(fd)

	file := os.NewFile(uintptr(fd), "raw-ethernet-frames")
	for {
		frame := make(EthernetFrame, 1024)
		_, err := file.Read(frame)
		if err != nil {
			log.Fatal(err)
		}

		packet := Packet(frame.Data())

		log.Println("Frame Data:")
		log.Printf("\tSource MAC Address: %v\n", frame.Source())
		log.Printf("\tDestination MAC Address: %v\n", frame.Destination())
		log.Println("Packet Data:")
		log.Printf("\tVersion: %v\n", packet.Version())
		log.Printf("\tIHL: %v\n", packet.IHL())
		log.Printf("\tProtocol: %v\n", packet.Protocol())
		log.Printf("\tSource IP: %v\n", packet.SourceIP())
		log.Printf("\tDestination IP: %v\n", packet.DestinationIP())
		if packet.Protocol() == 17 {
			udpDatagram := UDPDatagram(packet.Options())
			log.Println("UDP Datagram:")
			log.Printf("\tSource port: %v", udpDatagram.SourcePort())
			log.Printf("\tDestination port %v", udpDatagram.DestinationPort())
			log.Printf("\tLength: %v", udpDatagram.Length())
			log.Printf("\tBody: %v", udpDatagram.Body())
		}
		log.Println("----------------------------------------------------")
	}
}
