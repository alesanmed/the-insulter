package config

import "github.com/spf13/viper"

func LoadConfig() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetEnvPrefix("insulter")
	viper.BindEnv("postgres_dsn")
	viper.BindEnv("video_folder")

	return
}
