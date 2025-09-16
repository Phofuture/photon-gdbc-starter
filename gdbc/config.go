package gdbc

import "github.com/Phofuture/photon-core-starter/configuration"

func init() {
	configuration.Register(&databaseConfig)
}

var databaseConfig Config

type ConnectData struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Connection struct {
		MaxIdleConns      int `yaml:"maxIdleConns"`
		MaxOpenConns      int `yaml:"maxOpenConns"`
		MaxLifetimeSecond int `yaml:"maxLifetimeSecond"`
	} `yaml:",inline"`
	Schema    string `yaml:"schema"`
	Driver    string `yaml:"driver"` // e.g., "mysql", "postgres", etc.
	IfPrimary bool   `yaml:"ifPrimary"`
}

type Config struct {
	Database struct {
		Primary     ConnectData            `yaml:",inline"`
		DataSources map[string]ConnectData `yaml:"dataSources"`
	} `yaml:"database"`
}
