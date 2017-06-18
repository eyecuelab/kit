package cfg

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func InitViper(appName string, filePath string) error {
	viper.SetConfigType("yaml")
	if filePath != "" {
		viper.SetConfigFile(filePath)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(fmt.Sprintf("$HOME/%s", appName))
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "reading config")
	}
	return nil
}
