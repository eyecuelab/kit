package config

import (
	"log"

	"github.com/spf13/viper"
)

//RequiredString returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredString(key string) string {
	if viper.Get(key) == nil {
		log.Fatalf("missing required configuration val: %s", key)
	}
	return viper.GetString(key)

}
