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

func NewIkeSA() *IkeSA {
	return &IkeSA{Name: 0, State: "IKE_SA_INIT"}
}

// type childSA struct {
// }
