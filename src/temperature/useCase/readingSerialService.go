package useCase

import (
	"bytes"
	"log"
	"strings"

	"github.com/mikepb/go-serial"
)

//port="/dev/tty",baudrate=115200
func readSerialService(port string, baudrate int, endMesage string) []byte {
	options := serial.RawOptions
	options.BitRate = baudrate
	p, err := options.Open(port)
	// result := make([]byte, 100)
	var buffer bytes.Buffer
	if err != nil {
		log.Panic(err)
	}

	defer p.Close()
	i := 0
	for i < 1000 {
		buf := make([]byte, 1)
		if _, err := p.Read(buf); err != nil {
			log.Panic(err)
			break
		} else {
			// log.Println((buf))
		}

		if i == 0 {
			if strings.Contains(string(buf), "{") {
				buffer.Write(buf)
				i = i + 1
			}

		} else {

			if strings.Contains(string(buf), endMesage) {
				break
			}
			buffer.Write(buf)
			i = i + 1
		}
	}
	return buffer.Bytes()

}
