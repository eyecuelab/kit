//Package coerce provides c-style coercion for numeric types and strings, with additional guarantees for
//preservation of information.
package coerce

import (
	"math"

	"github.com/eyecuelab/kit/errorlib"
	"github.com/eyecuelab/kit/log"
)

const (
	maxConsecutiveIntInDoublePrecision                      = 2 << 53
	ErrOverflowInt64                   errorlib.ErrorString = "overflow of int64"
	ErrImpreciseConversion             errorlib.ErrorString = "data loss on conversion to float64"
	ErrWrongType                       errorlib.ErrorString = "could not coerce type"
)

var (
	allowOverFlow             = false
	allowImpreciseConversion  = true
	warnOnImpreciseConversion = true
	warnOnOverFlow            = true
)

func AllowOverflow(b bool)             { allowOverFlow = b }
func AllowImpreciseConversion(b bool)  { allowImpreciseConversion = b }
func WarnOnImpreciseConversion(b bool) { warnOnImpreciseConversion = b }

//Float64 converts a numeric type to a float64, if possible.
//Note for very large integers, this may lose information!
func Float64(v interface{}) (float64, error) {
	switch v := v.(type) {
	default:
		return 0, ErrWrongType
	case int:
		if int(float64(v)) != v {
			if warnOnImpreciseConversion {
				log.Warn(ErrImpreciseConversion)
			}
			if !allowImpreciseConversion {
				return float64(v), ErrImpreciseConversion
			}
		}
		return float64(v), nil

	case uint:
		if uint(float64(v)) != v {
			if warnOnImpreciseConversion {
				log.Warn(ErrImpreciseConversion)
			}
			if !allowImpreciseConversion {
				return float64(v), ErrImpreciseConversion
			}
		}
		return float64(v), nil

	case uint64:
		if uint64(float64(v)) != v {
			if warnOnImpreciseConversion {
				log.Warn(ErrImpreciseConversion)
			}
			if !allowImpreciseConversion {
				return float64(v), ErrImpreciseConversion
			}
		}
		return float64(v), nil

	case int64:
		if int64(float64(v)) != v {
			if warnOnImpreciseConversion {
				log.Warn(ErrImpreciseConversion)
			}
			if !allowImpreciseConversion {
				return float64(v), ErrImpreciseConversion
			}
		}
		return float64(v), nil

	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	}
}

//Int64 converts an integer type to an int64, if possible. Note that for very large unsigned ints,
//this may give an incorrect result!
func Int64(v interface{}) (int64, error) {
	switch v := v.(type) {
	default:
		return 0, ErrWrongType

	case uint:
		if v > math.MaxInt64 {
			if !allowOverFlow {
				return int64(v), ErrOverflowInt64
			} else if warnOnOverFlow {
				log.Warn(ErrOverflowInt64)
			}
		}
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			if !allowOverFlow {
				return int64(v), ErrOverflowInt64
			} else if warnOnOverFlow {
				log.Warn(ErrOverflowInt64)
			}
		}
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	}
}
