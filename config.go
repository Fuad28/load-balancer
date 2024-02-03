package main

type ConfigLoader interface {
	LoadConfig()
}

type CommonConfig struct {
	Port              int    `json:"port"`
	LBHealthCheckPath string `json:"lbHealthCheckPath"`
	Env               string `json:"env"` // DEV or PROD
}

type DevConfig struct {
	NoOfServers     int  `json:"noOfServers"`
	RandomServerOff bool `json:"randomServerOff"`
}

func (c *DevConfig) LoadConfig() error {
	// load config
	return nil
}

type ProdServer struct {
	Address         string `json:"address"` // host:port or domain
	HealthCheckPath string `json:"healthCheckPath"`
}
type ProdConfig struct {
	Servers         []ProdServer `json:"servers"`
	HealthCheckPath string       `json:"healthCheckPath"`
}

func (c *ProdConfig) LoadConfig() error {
	// load config
	return nil
}

type LoadBalancerConfig struct {
	CommonConfig
	DevConfig
	ProdConfig
}
