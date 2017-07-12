package cfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func CheckEnv(wanted ...string) error {
	var missing []string
	for _, s := range wanted {
		if ok := BeenSet(viper.Get(s)); ok == false {
			missing = append(missing, s)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("mising the following environment variable(s):\n\t%s", strings.Join(missing, "\n\t"))
	}
	return nil
}

func BeenSet(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return v.(int) != 0
	case string, []rune, []byte:
		return v.(string) != ""
	case []int, []int8, []int16, []int64, []uint, []uint16, []uint64, []float32, []float64:
		return v.([]interface{}) != nil
	case bool:
		return true // no way to tell
	default:
		var other interface{}
		return v.(interface{}) == other
	}
}
