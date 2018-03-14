package config

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/eyecuelab/kit/assets"
	"github.com/spf13/viper"
)

func Load(envPrefix string, configPath string) error {
	viper.SetConfigType("yaml")

	if len(configPath) > 0 {
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	} else {
		if data, err := assets.Get("data/bin/config.yaml"); err != nil {
			return err
		} else {
			viper.ReadConfig(bytes.NewBuffer(data))
		}
	}

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	for _, envVar := range viper.GetStringSlice("env") {
		if err := viper.BindEnv(envVar); err != nil {
			return err
		}
		if !viper.IsSet(envVar) {
			return errors.New(fmt.Sprintf("Env var is not set: %s", envVar))
		}
	}

	return nil
}
