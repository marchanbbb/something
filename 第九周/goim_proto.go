package main

import (
	"fmt"
	"net"
	"encoding/binary"
)

type ProtoStr struct {
	PKGLen int32
	HEADLen int16
	VER int16
	OPERATION int32
	SEQ int32
	BODY []byte
}

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	var data []byte
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	protoStr := &ProtoStr{
		PKGLen: binary.BigEndian.Int32(data[:2]),
		HEADLen: binary.BigEndian.Int32(data[2:3]),
		VER: binary.BigEndian.Int16(data[3:4]),
		OPERATION: binary.BigEndian.Int32(data[4:6]),
		SEQ: binary.BigEndian.Int32(data[6:8]),
		BODY: data[8:],
	}
	_ = protoStr
}