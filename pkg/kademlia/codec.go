package kademlia

import (
	"io"
	"bytes"
	"encoding/gob"
	"encoding/binary"
)

type messageCodec interface {
	deserialize(io.Reader) (*message, error)
	serialize(*message) ([]byte, error)
}

type gobCodec struct {}

func (codec *gobCodec) deserialize(conn io.Reader) (*message, error) {
	lengthBytes := make([]byte, 8)
	_, err := conn.Read(lengthBytes)
	if err != nil {
		return nil, err
	}

	lengthReader := bytes.NewBuffer(lengthBytes)
	length, err := binary.ReadUvarint(lengthReader)
	if err != nil {
		return nil, err
	}

	msgBytes := make([]byte, length)
	_, err = conn.Read(msgBytes)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewBuffer(msgBytes)
	msg := &message{}
	dec := gob.NewDecoder(reader)

	err = dec.Decode(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (codec *gobCodec) serialize(q *message) ([]byte, error) {
	var msgBuffer bytes.Buffer
	enc := gob.NewEncoder(&msgBuffer)
	err := enc.Encode(q)
	if err != nil {
		return nil, err
	}

	length := msgBuffer.Len()

	var lengthBytes [8]byte
	binary.PutUvarint(lengthBytes[:], uint64(length))

	var result []byte
	result = append(result, lengthBytes[:]...)
	result = append(result, msgBuffer.Bytes()...)

	return result, nil
}
