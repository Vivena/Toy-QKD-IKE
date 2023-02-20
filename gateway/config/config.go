package config

import (
	"time"

	viper "github.com/spf13/viper"

	"github.com/Vivena/Toy-QKD-IKE/gateway/crypto"
)

var Conf Config

var (
	// default connection timeout
	DefaultConnectionTimeout = 5 * time.Second
)

type Config struct {
	DefaultConnectionTimeout time.Duration

	// couchebase info
	QKD_Conf QKD
}

func (c *Config) load() error {
	// TODO: allow for the use of a config file and flags
	// TODO: proper handeling of the timeout with viper.GetDuration
	c.DefaultConnectionTimeout = 5 * time.Second

	viper.SetDefault("qkd.url", "127.0.0.1")
	viper.SetDefault("qkd.port", "8000")
	viper.SetDefault("qkd.saeID", "test")

	qkdURL := viper.GetString("qkd.url")
	qkdPort := viper.GetString("qkd.port")
	qkdSaeID := viper.GetString("qkd.saeID")
	c.QKD_Conf = crypto.NewQKD(qkdURL, qkdPort, qkdSaeID)

	return nil
}

func GlobalConfig() (*Config, error) {
	c := &Config{}
	if err := c.load(); err != nil {
		return nil, err
	}
	return c, nil
}
