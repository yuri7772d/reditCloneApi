package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server
		OAuth2   *OAuth2
		State    *State
		Database *Database
	}
	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
	}

	OAuth2 struct {
		PlayerRedirectUrl string `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientId          string `mapstructure:"clientId" validate:"required"`
		ClientSecret      string `mapstructure:"clientSecret" validate:"required"`
		Endpoints         endpoints
		Scopes            []string `mapstructure:"scopes" validate:"required"`
		UserInfoUrl       string   `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl         string   `mapstructure:"revokeUrl" validate:"required"`
	}

	endpoints struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}
	State struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAt time.Duration `mapstructure:"expiresAt" validate:"required"`
		Issuer    string        `mapstructure:"issuer" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Dbname   string `mapstructure:"dbname" validate:"required"`
		Sslmode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

var (
	instance *Config
	once     sync.Once
)

func Get() (*Config, error) {
	once.Do(func() {

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		instance = &Config{}
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&instance); err != nil {
			fmt.Printf("Unmarshal error:\n%+v\n", err)
			panic(err)
		}
		validating := validator.New()
		if err := validating.Struct(instance); err != nil {
			panic(err)
		}

	})

	return instance, nil
}
