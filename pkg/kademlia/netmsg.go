package kademlia

import (
	"encoding/gob"
)

const (
	messageTypePing = iota
	messageTypeStore
	messageTypeFindNode
	messageTypeFindValue
)

type message struct {
	Sender     *NetworkNode
	Receiver   *NetworkNode
	ID         int64
	Error      error
	Type       int
	IsResponse bool
	Data       interface{}
}

type queryDataFindNode struct {
	Target []byte
}

type queryDataFindValue struct {
	Target []byte
}

type queryDataStore struct {
	Data       []byte
	Publishing bool // Whether or not we are the original publisher
}

type responseDataFindNode struct {
	Closest []*NetworkNode
}

type responseDataFindValue struct {
	Closest []*NetworkNode
	Value   []byte
}

type responseDataStore struct {
	Success bool
}

func netMsgInit() {
	gob.Register(&queryDataFindNode{})
	gob.Register(&queryDataFindValue{})
	gob.Register(&queryDataStore{})
	gob.Register(&responseDataFindNode{})
	gob.Register(&responseDataFindValue{})
	gob.Register(&responseDataStore{})
}
