package config

import (
	"log"

	"github.com/spf13/viper"
)

//RequiredString returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredString(key string) string {
	FatalCheck(key)
	return viper.GetString(key)
}

//RequiredInt  returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredInt(key string) int {
	FatalCheck(key)
	return viper.GetInt(key)
}

//RequiredFloat64 returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredFloat64(key string) float64 {
	FatalCheck(key)
	return viper.GetFloat64(key)
}

//RequiredStringSlice returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredStringSlice(key string) []string {
	FatalCheck(key)
	return viper.GetStringSlice(key)
}

//FatalCheck checks that a key exists in the viper configuration. If not, it calls fatalf.
func FatalCheck(key string) {
	if viper.Get(key) == nil {
		log.Fatalf("missing required configuration val: %s", key)
	}
}
