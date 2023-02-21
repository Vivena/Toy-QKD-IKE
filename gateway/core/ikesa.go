package core

import "sync"

type IkeSA struct {
	SALock sync.Mutex
	Name   uint32
	State  string
	Key_ID string
	Key    []byte
	HasKey bool
}

// NewIkeSA: create the SA with the state "IKE_SA_INIT", and set the name to 0 while we wait for
// the responder SPI
func NewIkeSA() *IkeSA {
	return &IkeSA{Name: 0, State: "IKE_SA_INIT"}
}

// TODO: create getters and setters for IkeSA
