package config

import (
	"github.com/spf13/viper"
)

func Load(envPrefix string, configPath string) error {
	viper.SetConfigType("yaml")

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		// viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	// # TODO: raise error if not present
	// for _, envVar := range viper.GetStringSlice("env") {
	// 	if err := viper.BindEnv(envVar); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
