package main

import "fmt"

type Config struct {
	server string
	port   int
}

func NewConfig() *Config {
	return &Config{
		server: "localhost",
		port:   80,
	}
}

func (config *Config) WithServer(server string) *Config {
	config.server = server
	return config
}

func (config *Config) WithPort(port int) *Config {
	config.port = port
	return config
}

func main() {
	config := NewConfig().
		WithServer("remotehost").
		WithPort(8080)

	fmt.Println(config)
}
