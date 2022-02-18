package config

import (
	"github.com/spf13/viper"
)

func Load(envPrefix string, configPath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(envPrefix)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	// viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	err = viper.MergeInConfig()
	if err != nil {
		return err
	}

	// for _, envVar := range viper.GetStringSlice("env") {
	// 	if err := viper.BindEnv(envVar); err != nil {
	// 		return err
	// 	}
	// 	if !viper.IsSet(envVar) {
	// 		return errors.New(fmt.Sprintf("Env var is not set: %s", envVar))
	// 	}
	// }

	return nil
}
