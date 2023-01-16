package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	ApiConfig
	DbConfig
}

type ApiConfig struct {
	Port string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// Método sempre chamado na inicializzação da aplicação
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {

	//Seta de qual arquivo vai pegar
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//Api
	cfg = new(config)
	cfg.ApiConfig = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	//DbConfig
	cfg.DbConfig = DbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	//
	return nil
}

func GetDB() DbConfig {
	return cfg.DbConfig
}

func GetServerPort() string {
	return cfg.ApiConfig.Port
}
