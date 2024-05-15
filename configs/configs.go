package configs

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/typesense/typesense-go/typesense/api"
)

type Config struct {
	TS struct {
		ApiKey string     `mapstructure:"TS_API_KEY"`
		Nodes  []api.Node `mapstructure:"TS_NODES"`
	}
	MultipleTS map[string]struct {
		ApiKey string     `mapstructure:"MULTIPLE_TS_API_KEY"`
		Nodes  []api.Node `mapstructure:"MULTIPLE_TS_NODES"`
	}
	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Port     string `mapstructure:"PORT"`
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		}
	}
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		err = viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to unmarshal config")
		}
	})

	return &cfg
}
