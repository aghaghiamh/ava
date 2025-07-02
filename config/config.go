package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/aghaghiamh/ava/handler/httpserver"
	"github.com/aghaghiamh/ava/handler/httpserver/userhandler"
	mysql "github.com/aghaghiamh/ava/repository"
)

const (
	ENV_PREFIX = "AVA_"
)

type Config struct {
	HandlerConfig userhandler.HandlerConfig `mapstructure:"handler_params"`
	DB            mysql.Config              `mapstructure:"db_params"`
	Server        httpserver.HttpConfig     `mapstructure:"server_params"`
}

func LoadConfig() Config {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	viper.SetDefault("server_params.graceful_shutdown_timeout", "5s")

	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(ENV_PREFIX)

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.BindEnv("db_params.username")
	viper.BindEnv("db_params.password")

	setEnvValues()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %w", err))
	}

	return config
}

func setEnvValues() {
	// Define the config keys that should be set from environment variables
	envKeys := []string{
		"db_params.username",
		"db_params.password",
	}

	for _, key := range envKeys {
		envKey := ENV_PREFIX + strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

		if val, exists := os.LookupEnv(envKey); exists {
			// Set the value in Viper so it appears in AllSettings()
			viper.Set(key, val)
		}
	}
}
