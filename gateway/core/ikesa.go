package core

type IkeSA struct {
	Name  uint32
	State string
}

func NewIkeSA() *IkeSA {
	return &IkeSA{Name: 0, State: "IKE_SA_INIT"}
}

// type childSA struct {
// }
