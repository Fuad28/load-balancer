package main

type ConfigLoader interface {
	LoadConfig()
}

type CommonConfig struct {
	port int
}

type DevConfig struct {
	CommonConfig
	noOfServers     int
	randomServerOff bool
}

type ProdConfig struct {
	CommonConfig
	serversAddress  []string
	host            string
	port            string
	healthCheckPath string
}

func (c *DevConfig) LoadConfig() error {
	// load config
	return nil
}

func (c *ProdConfig) LoadConfig() error {
	// load config
	return nil
}

type LoadBalancerConfig struct {
	ConfigLoader ConfigLoader
}
