package model

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Database string
	User     string
	Password string
	Host     string
}

var config *Config

func (c *Config) GetDBConnectionString() string {
	var b strings.Builder
	fmt.Fprintf(&b, "postgresql://%s:%s@%s:5432/%s?sslmode=disable", c.User, c.Password, c.Host, c.Database)

	return b.String()
}

func (c *Config) Init() {
	if database, ok := os.LookupEnv("TC_DATABASE"); ok != true {
		c.Database = "test"
	} else {
		c.Database = database
	}

	if user, ok := os.LookupEnv("TC_USER"); ok != true {
		c.User = "test"
	} else {
		c.User = user
	}

	if password, ok := os.LookupEnv("TC_PASSWORD"); ok != true {
		c.Password = "test"
	} else {
		c.Password = password
	}

	if host, ok := os.LookupEnv("TC_HOST"); ok != true {
		c.Host = "localhost"
	} else {
		c.Host = host
	}
}

func SetConfig(cfg *Config) {
	config = cfg
}

func GetConfig() *Config {
	return config
}
