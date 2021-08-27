package main

type Config struct {
	bindAddr    string  `toml:"bind_addr"`
	mocksFolder string  `toml:"mocks_folder"`
	Routes      []Route `toml:"routes"`
}

type Route struct {
	Path     string `toml:"path"`
	Filename string `toml:"filename"`
}

func NewConfig() *Config {
	return &Config{
		bindAddr:    ":8080",
		mocksFolder: "mocks",
		Routes:      []Route{},
	}
}
