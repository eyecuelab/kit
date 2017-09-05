package config

import (
	"log"
	"os"

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

//RequiredSlice returns the value of the configuration variable specified by key.
//If that configuration var does not exist, it calls log.Fatalf
func RequiredSlice(key string) []interface{} {
	FatalCheck(key)
	return viper.Get(key).([]interface{})
}

//FatalCheck checks that a key exists in the viper configuration. If not, it calls fatalf.
func FatalCheck(key string) {
	if viper.Get(key) == nil {
		log.Fatalf("missing required configuration val: %s", key)
	}
}

//RequiredEnv checks that the specified environment variable exists. If not, it calls log.Fatalf.
func RequiredEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	log.Fatalf("missing required environment variable %s", key)
	return ""
}
