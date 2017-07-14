package env

//kit.env contains tools to help with the assignment of environment variables
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

//Key is the key string of an environment variable.
type Key string

//Value is an alias for os.GetEnv(string(key))
func (key Key) Value() string {
	return os.Getenv(string(key))
}

//Lookup is an alias for os.LookupEnv(string(key))
func (key Key) Lookup() (string, bool) {
	return os.LookupEnv(string(key))
}

//SetString sets the string address pointed to by sp to key.Value().
func (key Key) SetString(sp *string) error {
	val, ok := key.Lookup()
	if !ok {
		return fmt.Errorf("missing environment variable %s", key)
	}
	*sp = val
	return nil
}

//SetInt sets the int address pointed to by ip to key.Value().
func (key Key) SetInt(ip *int) error {
	val, ok := key.Lookup()
	if !ok {
		return fmt.Errorf("env.Key.SetInt: missing environment variable %s", key)
	}
	asInt, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("env.Key.SetInt: strconv.Atoi: %v", err)
	}
	*ip = asInt
	return nil
}

//SetBool sets the boolean address pointed to by bp to key.Value()
func (key Key) SetBool(bp *bool) error {
	str := string(key)
	val, ok := os.LookupEnv(str)
	if !ok {
		return fmt.Errorf("env.Key.SetBool: missing environment variable %s", key)
	}
	asBool, err := strconv.ParseBool(val)
	if err != nil {
		return fmt.Errorf("env.Key.SetInt: strconv.ParseBool: %v", err)
	}
	*bp = asBool
	return nil
}

//Check checks the local environment to see that all of the environment variables specified in config are set.
func Check() {
	var missing []string
	for _, key := range viper.GetStringSlice("env_var_names") {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("CheckEnv(): missing the following environmental variables: \n\t%s", strings.Join(missing, "\n\t"))
	}
}
