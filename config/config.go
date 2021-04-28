package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Mode string `yaml:"mode"`
	Port uint16 `yaml:"port"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
	Loc      string `yaml:"loc"`
}

// Read the config file, and returns the server and database config.
func Read(filename string) (server Server, db Database) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error occurred in reading config file.")
	}

	var root struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
	}
	err = yaml.Unmarshal(bytes, &root)
	if err != nil {
		panic("Config file malformed.")
	}
	return root.Server, root.Database
}
