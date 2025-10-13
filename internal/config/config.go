package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Session struct {
		MaxAge   int  `mapstructure:"maxAge"`
		IsProd   bool `mapstructure:"isProd"`
		HttpOnly bool `mapstructure:"httpOnly"`
	} `mapstructure:"session"`

	Auth0 struct {
		ClientID     string `mapstructure:"clientId"`
		ClientSecret string `mapstructure:"clientSecret"`
		Domain       string `mapstructure:"domain"`
		CallbackURL  string `mapstructure:"callbackUrl"`
	} `mapstructure:"auth0"`

	Google struct {
		GoogleKey    string `mapstructure:"googleKey"`
		GoogleSecret string `mapstructure:"googleSecret"`
		CallbackURL  string `mapstructure:"callbackUrl"`
	} `mapstructure:"google"`

	AzureAD struct {
		AzureADKey    string `mapstructure:"azureadKey"`
		AzureADSecret string `mapstructure:"azureadSecret"`
		CallbackURL   string `mapstructure:"callbackUrl"`
	} `mapstructure:"azuread"`

	Saml struct {
		IDPMetadataURL string `mapstructure:"idpMetadataURL"`
		RootURL        string `mapstructure:"rootURL"`
		EntityID       string `mapstructure:"entityID"`
	} `mapstructure:"saml"`

	Cors struct {
		AllowedOrigins   []string `mapstructure:"allowedOrigins"`
		AllowedMethods   []string `mapstructure:"allowedMethods"`
		AllowedHeaders   []string `mapstructure:"allowedHeaders"`
		AllowCredentials bool     `mapstructure:"allowCredentials"`
		MaxAge           int      `mapstructure:"maxAge"`
	} `mapstructure:"cors"`
}

func LoadConfig() (*Config, error) {

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	return loadConfig(env)
}

func LoadConfigForEnv(env string) (*Config, error) {
	return loadConfig(env)
}

func loadConfig(env string) (*Config, error) {

	viper.SetConfigType("yaml")
	var configFile string

	switch env {
	case "dev", "development":
		configFile = "internal/config/config.dev.yaml"
	case "prod", "production":
		configFile = "internal/config/config.prod.yaml"
	case "staging":
		configFile = "internal/config/config.staging.yaml"
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file %s: %s", configFile, err)
	}

	// Map the config file values into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %v", err)
	}

	return &config, nil
}
