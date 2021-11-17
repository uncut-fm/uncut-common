package config

type DBConfigs struct {
	Host                  string `json:"host"`
	Port                  string `json:"port"`
	DBName                string `json:"db_name"`
	User                  string `json:"user"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `yaml:"max_open_conns"`
	MaxIdleConnections    int    `yaml:"max_idle_conns"`
	ConnectionMaxLifetime string `yaml:"conn_max_lifetime"`
}
