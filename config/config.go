package config

import "time"

var cfg *Config

type (
	Config struct {
		App     `yaml:"app"`
		GRPC    `yaml:"grpc"`
		Log     `yaml:"logger"`
		MongoDB `yaml:"mongodb"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	GRPC struct {
		Port string `env-required:"true" yaml:"port" env:"GRPC_PORT"`
	}

	MongoDB struct {
		URL                string           `env-required:"true" yaml:"url"`
		MongoDBConnTimeout time.Duration    `yaml:"connTimeout" default:"10s"`
		MongoDatabase      string           `env-required:"true" yaml:"database"`
		MongoCollections   MongoCollections `env-required:"true" yaml:"collections"`
	}
	MongoCollections struct {
		Order `env-required:"true" yaml:"order"`
	}

	Order struct {
		Collection string        `env-required:"true" yaml:"collection"`
		MaxTime    time.Duration `yaml:"max-time" default:"2s"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
