package config

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Grpc struct {
		Port                string `mapstructure:"port"`
		UsersServiceAddress string `mapstructure:"users_service_address"`
	} `mapstructure:"grpc"`

	Logging struct {
		IsProduction bool   `mapstructure:"isProduction"`
		VectorURL    string `mapstructure:"vectorURL"`
	} `mapstructure:"logging"`

	Postgres struct {
		User     string `mapstructure:"user"`
		Name     string `mapstructure:"name"`
		Password string `mapstructure:"password"`
		Address  string `mapstructure:"address"`
	} `mapstructure:"postgres"`

	RabbitMQ struct {
		Address    string `mapstructure:"address"`
		EmailQueue string `mapstructure:"email_queue"`
	} `mapstructure:"rabbitmq"`

	Redis struct {
		TokensTTL time.Duration `mapstructure:"tokens_ttl"`
		Address   string        `mapstructure:"address"`
		Password  string        `mapstructure:"password"`
		DB        int           `mapstructure:"db"`
	} `mapstructure:"redis"`
}

var Cfg Config

func LoadConfig(path string) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config not found, used env and default values")
		} else {
			log.Error().Msg("Failed read config file")
		}
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			envValue := os.Getenv(envVar)
			if envValue != "" {
				v.Set(k, envValue)
			}
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Error().Msg("Failed unmarshal config")
	}

	if err := validateConfig(&cfg); err != nil {
		log.Fatal().Msg("Not valid config")
	}

	Cfg = cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("grpc.port", "8000")

	v.SetDefault("postgres.address", "http://localhost:5432")
	v.SetDefault("postgres.user", "admin")
	v.SetDefault("postgres.password", "password")
	v.SetDefault("postgres.name", "db")

	v.SetDefault("logging.isProduction", false)
	v.SetDefault("logging.vectorURL", "http://vector:9880")

	v.SetDefault("rabbitmq.address", "amqp://guest:guest@localhost:5672/")
	v.SetDefault("rabbitmq.notifications_queue", "tasks")
	v.SetDefault("rabbitmq.email_queue", "tasks")
}

func validateConfig(cfg *Config) error {
	return nil
}
