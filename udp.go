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
		if packet.Protocol() != 17 {
			continue
		}

		udpDatagram := UDPDatagram(packet.Options())
		log.Printf("%v", packet.IHL())
		log.Printf("received message: %s", udpDatagram.Body())
	}
}
