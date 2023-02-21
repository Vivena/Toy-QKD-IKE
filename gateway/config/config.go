package config

import (
	"time"

	viper "github.com/spf13/viper"

	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
)

// TODO: Need to finish this part at a later date
// this is where we read a conf file and create a config for our instance
// the code bellos is not yet usable, nor is it incorporated in the rest of the implementation

var (
	// default connection timeout
	DefaultConnectionTimeout = 5 * time.Second
)

type Config struct {
	DefaultConnectionTimeout time.Duration

	QKD_Conf crypto.QKD
}

func (c *Config) load() error {
	// TODO: allow for the use of a config file and flags
	// TODO: proper handeling of the timeout with viper.GetDuration
	c.DefaultConnectionTimeout = 5 * time.Second

	viper.SetDefault("qkd.url", "127.0.0.1")
	viper.SetDefault("qkd.port", "8000")
	viper.SetDefault("qkd.saeID", 1)

	qkdURL := viper.GetString("qkd.url")
	qkdPort := viper.GetString("qkd.port")
	qkdSaeID := viper.GetUint16("qkd.saeID")
	c.QKD_Conf = *crypto.NewQKD(qkdURL, qkdPort, qkdSaeID)

	return nil
}

func GlobalConfig() (*Config, error) {
	c := &Config{}
	if err := c.load(); err != nil {
		return nil, err
	}
	return c, nil
}
