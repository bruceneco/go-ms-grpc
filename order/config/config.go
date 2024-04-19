package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	EnvMode              = "env"
	EnvAppPort           = "app_port"
	EnvDBUrl             = "db_url"
	EnvPaymentServiceURL = "payment_service.url"
)

type Config struct {
	viper *viper.Viper
}

func NewConfig() *Config {
	v := viper.New()
	v.AddConfigPath("..")
	v.SetConfigFile("env.yml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	cfg := &Config{v}
	cfg.setDefaults()
	return cfg
}

func (c *Config) setDefaults() {
	c.viper.SetDefault("env", "development")
	c.viper.SetDefault("app_port", "8080")
}

func (c *Config) GetString(name string) string {
	return c.viper.GetString(name)
}
