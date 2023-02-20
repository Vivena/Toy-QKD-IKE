package core

type IkeSA struct {
	Name            string             `vici:"-"`
	State           string             `vici:"state"`
	LocalVirtualIPs []string           `vici:"local-vips"`
	ChildSAs        map[string]childSA `vici:"child-sas"`
}

type childSA struct {
	Name  string `vici:"name"`
	State string `vici:"state"`
}
