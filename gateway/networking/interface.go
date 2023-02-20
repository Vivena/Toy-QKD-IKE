type Interface struct {
}

func (i *Interface) Execute() {
	conf, err := config.GlobalConfig()
	if err != nil {
		return err
	}
}