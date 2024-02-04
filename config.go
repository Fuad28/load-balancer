package main

import (
	"strings"

	validator "github.com/asaskevich/govalidator"
)

const CONFIGPATH = "config.json"

type ServerConf struct {
	Address         string `json:"address" valid:"required"` // host:port or domain
	HealthCheckPath string `json:"healthCheckPath" valid:"required"`
}

type LoadBalancerConfig struct {
	Port    int          `json:"port" valid:"required"`
	Env     string       `json:"env" valid:"required"` // DEV or PROD
	Servers []ServerConf `json:"servers" valid:"required"`

	// dev configs
	NoOfServers     int  `json:"noOfServers"`
	RandomServerOff bool `json:"randomServerOff"`
}

func (lbConfig *LoadBalancerConfig) LoadConfig() error {
	err := LoadFile[LoadBalancerConfig](CONFIGPATH, lbConfig)

	OnErrorPanic(err, "Error loading config")

	_, err = validator.ValidateStruct(lbConfig)

	OnErrorPanic(err, "config validation error")

	if (strings.ToLower(config.Env) == "dev") && (lbConfig.NoOfServers <= 0) {
		OnErrorPanic(err, "NoOfServers has to be greater than 1")
	}

	return nil
}
