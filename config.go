package main

type ConfigLoader interface {
	LoadConfig()
}

type CommonConfig struct {
	Port            int
	healthCheckPath string
}

type DevConfig struct {
	CommonConfig
	noOfServers     int
	randomServerOff bool
}

func (c *DevConfig) LoadConfig() error {
	// load config
	return nil
}

type ProdServer struct {
	Address         string // host:port or domain
	healthCheckPath string
}
type ProdConfig struct {
	CommonConfig
	Servers         []ProdServer
	healthCheckPath string
}

func (c *ProdConfig) LoadConfig() error {
	// load config
	return nil
}

type LoadBalancerConfig struct {
	ConfigLoader ConfigLoader
}
