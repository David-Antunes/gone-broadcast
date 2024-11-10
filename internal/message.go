package internal

import (
	"bytes"
	"encoding/gob"
	"net"
)

type Msg struct {
	Ip   []byte
	Port int
}

func SendMsg(ip net.IP, port int) *bytes.Buffer {

	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	enc.Encode(Msg{
		Ip:   ip,
		Port: port,
	})

	return buf
}

func ReceiveMsg(data *bytes.Buffer) (Msg, error) {

	var msg Msg

	dec := gob.NewDecoder(data)

	err := dec.Decode(&msg)

	if err != nil {
		return Msg{}, err
	}

	return msg, nil
}
