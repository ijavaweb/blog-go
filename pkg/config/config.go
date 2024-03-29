package config

import "os"

func LoadConfig() string {
	_port, present := os.LookupEnv("PORT")
	if !present {
		return "127.0.0.1:8083"
	}

	return "127.0.0.1" + ":" + _port
}
