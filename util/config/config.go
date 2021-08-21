package config

import (
	"sync"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/spf13/viper"
)

var (
	config *EnvironmentVariables
	once   sync.Once
)

// GetConfigEnvironment to read initial config
func GetConfigEnvironment() *EnvironmentVariables {
	once.Do(func() {

		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			logger.Error("Error to read configs: ", err)
			panic(err)
		}

		config = &EnvironmentVariables{}
		config.MySQL.Username = viper.GetString("DB_USER")
		config.MySQL.Password = viper.GetString("DB_PASSWORD")
		config.MySQL.Host = viper.GetString("DB_HOST")
		config.MySQL.Port = viper.GetString("DB_PORT")
		config.MySQL.DBName = viper.GetString("DB_NAME")
		config.MySQL.MaxLifeInMinutes = viper.GetInt("MAX_LIFE_IN_MINUTES")
		config.MySQL.MaxIdleConns = viper.GetInt("MAX_IDLE_CONNS")
		config.MySQL.MaxOpenConns = viper.GetInt("MAX_OPEN_CONNS")
	})

	return config
}

// EnvironmentVariables is environment variables configs
type EnvironmentVariables struct {
	MySQL mysqlConfig
}

type mysqlConfig struct {
	Username         string
	Password         string
	Host             string
	Port             string
	DBName           string
	MaxLifeInMinutes int
	MaxIdleConns     int
	MaxOpenConns     int
}
