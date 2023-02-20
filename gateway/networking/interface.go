package networking

import "github.com/Vivena/Toy-QKD-IKE/gateway/config"

type Interface struct {
	
}

func (i *Interface) NewInterface() error {
	_, err := config.GlobalConfig()
	if err != nil {
		return err
	}
	return nil
}

func (i *Interface) Start() error {
	_, err := config.GlobalConfig()
	if err != nil {
		return err
	}
	return nil
}
