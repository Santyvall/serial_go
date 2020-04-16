package main

import (
	"C"
	"log"

	"github.com/mikepb/go-serial"
)
import "fmt"

func main() {
	options := serial.RawOptions
	options.BitRate = 115200
	options.Parity = serial.PARITY_MARK

	var COININ = []byte{0x10, 0x11, 0x12, 0x13, 0x14}

	p, err := options.Open("/dev/ttyUSB0")
	if err != nil {
		log.Panic(err)
	}

	defer p.Close()

	n, err := p.Write(COININ)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Write : ", n)

	buf := make([]byte, 10)
	if _, err := p.Read(buf); err != nil {
		log.Panic(err)
	} else {
		log.Println(buf)
	}
}
