package configs

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config is a struct that will receive configuration options via environment
// variables.
type Config struct {
	App struct {
		CORS struct {
			AllowCredentials bool     `mapstructure:"ALLOW_CREDENTIALS"`
			AllowedHeaders   []string `mapstructure:"ALLOWED_HEADERS"`
			AllowedMethods   []string `mapstructure:"ALLOWED_METHODS"`
			AllowedOrigins   []string `mapstructure:"ALLOWED_ORIGINS"`
			Enable           bool     `mapstructure:"ENABLE"`
			MaxAgeSeconds    int      `mapstructure:"MAX_AGE_SECONDS"`
		}
		Env  string `mapstructure:"ENV"`
		Http struct {
			Port string `mapstructure:"PORT"`
		} `mapstructure:"HTTP"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Name     string `mapstructure:"NAME"`
		URL      string `mapstructure:"URL"`
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		}
	} `mapstructure:"APP"`
	Database struct {
		Read struct {
			Host        string `mapstructure:"HOST"`
			Port        int    `mapstructure:"PORT"`
			Name        string `mapstructure:"NAME"`
			User        string `mapstructure:"USER"`
			Password    string `mapstructure:"PASSWORD"`
			SslMode     string `mapstructure:"SSLMODE"`
			MaxIdleConn int    `mapstructure:"MAX_IDLE_CONN"`
			MaxOpenConn int    `mapstructure:"MAX_OPEN_CONN"`
		} `mapstructure:"Read"`
		Write struct {
			Host        string `mapstructure:"HOST"`
			Port        int    `mapstructure:"PORT"`
			Name        string `mapstructure:"NAME"`
			User        string `mapstructure:"USER"`
			Password    string `mapstructure:"PASSWORD"`
			SslMode     string `mapstructure:"SSLMODE"`
			MaxIdleConn int    `mapstructure:"MAX_IDLE_CONN"`
			MaxOpenConn int    `mapstructure:"MAX_OPEN_CONN"`
		} `mapstructure:"WRITE"`
	} `mapstructure:"DATABASE"`
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	return &conf
}
