package main

import (
	"errors"
	"strings"

	validator "github.com/asaskevich/govalidator"
)

const CONFIGPATH = "config.json"

type ServerConf struct {
	Address         string `json:"address" valid:"required"` // host:port or domain
	HealthCheckPath string `json:"healthCheckPath" valid:"required"`
}

type LoadBalancerConfig struct {
	Port int    `json:"port" valid:"required"`
	Env  string `json:"env" valid:"required"` // DEV or PROD

	// dev specific configs
	NoOfServers     int  `json:"numberOfServers"`
	RandomServerOff bool `json:"randomServerOff"`

	// prod specific configs
	Servers []ServerConf `json:"servers"`
}

func (lbConfig *LoadBalancerConfig) LoadConfig() error {
	err := LoadFile[LoadBalancerConfig](CONFIGPATH, lbConfig)

	OnErrorPanic(err, "Error loading config")

	_, err = validator.ValidateStruct(lbConfig)

	OnErrorPanic(err, "config validation error")

	if (strings.ToLower(config.Env) == "dev") && (lbConfig.NoOfServers <= 0) {
		OnErrorPanic(errors.New("NoOfServers has to be greater than 0"), "")
	}

	if (strings.ToLower(config.Env) == "prod") && (len(lbConfig.Servers) == 0) {
		OnErrorPanic(errors.New("at least one server is required to redirect traffic"), "")
	}

	return nil
}
