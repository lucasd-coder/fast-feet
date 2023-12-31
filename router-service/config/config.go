package config

import "time"

var cfg *Config

type (
	Config struct {
		App         `yaml:"app"`
		Server      `yaml:"server"`
		Log         `yaml:"logger"`
		Integration `yaml:"integration"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		AesKey  string `env-required:"true" yaml:"aes-key" env:"AES_KEY"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Integration struct {
		RabbitMQ      `env-required:"true" yaml:"rabbit-mq"`
		GrpcClient    `env-required:"true" yaml:"grpc"`
		OpenTelemetry `env-required:"true" yaml:"otlp"`
	}

	RabbitMQ struct {
		Topic `env-required:"true" yaml:"topic"`
	}

	Topic struct {
		TopicUserEvents  `env-required:"true" yaml:"user-events"`
		TopicOrderEvents `env-required:"true" yaml:"order-events"`
	}

	Server struct {
		Port         string        `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		ReadTimeout  time.Duration `yaml:"readTimeout" default:"10s"`
		WriteTimeout time.Duration `yaml:"writeTimeout" default:"10s"`
	}

	TopicUserEvents struct {
		URL         string        `env-required:"true" yaml:"url" env:"USER_EVENTS_URL"`
		MaxRetries  int           `yaml:"max-retries" env-default:"3"`
		WaitingTime time.Duration `yaml:"waiting-time" env-default:"2s"`
	}

	TopicOrderEvents struct {
		URL         string        `env-required:"true" yaml:"url" env:"ORDER_EVENTS_URL"`
		MaxRetries  int           `yaml:"max-retries" env-default:"3"`
		WaitingTime time.Duration `yaml:"waiting-time" env-default:"2s"`
	}

	GrpcClient struct {
		BusinessService `env-required:"true" yaml:"business-service"`
	}

	BusinessService struct {
		URL      string `env-required:"true" yaml:"url" env:"BUSINESS_SERVICE_URL"`
		MaxRetry uint   `env-required:"true" yaml:"max-retry"`
	}

	OpenTelemetry struct {
		URL      string        `env-required:"true" yaml:"url" env:"OTEL_EXPORTER_OTLP_ENDPOINT"`
		Protocol string        `env-required:"true" yaml:"protocol" env:"OTEL_EXPORTER_OTLP_PROTOCOL"`
		Timeout  time.Duration `env-required:"true" yaml:"timeout" env:"OTEL_EXPORTER_OTLP_TIMEOUT"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
