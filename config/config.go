package config

import (
	"strings"

	"github.com/spf13/viper"
)

// viper config struct
type (
	ViperConfig struct {
		DatasourceConfig DatasourceConfig  `mapstructure:"datasource"`
		ServerConfig     ViperServerConfig `mapstructure:"server"`
		Observability    Observability     `mapstructure:"observability"`
	}
	// DatasourceConfig is the datasource config.
	DatasourceConfig struct {
		Mysql  MysqlConfig  `mapstructure:"mysql"`
		Redis  RedisConfig  `mapstructure:"redis"`
		Sqlite SqliteConfig `mapstructure:"sqlite"`
	}

	// viper server config struct
	ViperServerConfig struct {
		Grpc GrpcConfig `mapstructure:"grpc"`
		Http HttpConfig `mapstructure:"http"`
	}

	Observability struct {
		Tracing Tracing `mapstructure:"tracing"`
		Logging Logging `mapstructure:"logging"`
	}

	Logging struct {
		Level  string `mapstructure:"level"`
		Enable bool   `mapstructure:"enable"`
	}

	Tracing struct {
		Enabled bool   `mapstructure:"enabled"`
		Zipkin  Zipkin `mapstructure:"zipkin"`
	}

	Zipkin struct {
		Url         string `mapstructure:"url"`
		ServiceName string `mapstructure:"service_name"`
	}

	HttpConfig struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}

	// grpc config struct
	GrpcConfig struct {
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		Production bool   `mapstructure:"production"`
	}

	// MysqlConfig is the mysql config.
	MysqlConfig struct {
		Enabled  bool   `mapstructure:"enabled"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	}

	// RedisConfig is the redis config.
	RedisConfig struct {
		Enabled  bool   `mapstructure:"enabled"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Database int    `mapstructure:"database"`
	}

	SqliteConfig struct {
		Enabled bool   `mapstructure:"enabled"`
		Dns     string `mapstructure:"dns"`
	}
)

// config is the config.
type Config struct {
	address string
}

// newConfig creates a new config.
func NewConfig(address string) (*Config, error) {
	return &Config{address: address}, nil
}

// load config from path
func (c *Config) Load(cfg interface{}) error {
	// create a new viper instance
	v := viper.New()
	// set the config name
	v.SetConfigFile(c.address)
	// set env replacement
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//
	v.AutomaticEnv()
	// read the config
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// unmarshal the config
	if err := v.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
