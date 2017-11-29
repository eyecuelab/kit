package web

import (
	"regexp"
	"strings"

	"github.com/labstack/echo"
	emw "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type (
	AuthedContextLookup interface {
		Lookup(echo.Context) (echo.Context, error)
		Context(echo.Context) echo.Context
	}
	// AuthedConfig config for Authed middleware.
	AuthedConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper emw.Skipper
	}
)

var (
	// DefaultAuthedConfig default Authed middleware config.
	DefaultAuthedConfig = AuthedConfig{
		Skipper: AuthedSkipper(),
	}
)

type authSkipperConfig map[string]*regexp.Regexp

func AuthedSkipper() func(echo.Context) bool {
	config := viper.GetStringMapString("skipjwt")

	if config == nil || len(config) == 0 {
		return emw.DefaultSkipper
	}

	skipper := authSkipperConfig{}
	for method, exp := range config {
		skipper[strings.ToUpper(method)] = regexp.MustCompile(exp)
	}

	return func(c echo.Context) bool {
		if c.Request().Method == echo.OPTIONS {
			return true
		}
		re, ok := skipper[c.Request().Method]
		if !ok {
			return false
		}

		if hasAuthHeader(c) {
			return false
		}

		return re.MatchString(c.Request().URL.Path)
	}
}

// AuthedWithConfig ...
func AuthedWithConfig(config AuthedConfig, cl AuthedContextLookup) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultAuthedConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(cl.Context(c))
			}
			ac, err := cl.Lookup(c)
			if err != nil {
				return err
			}

			return next(ac)
		}
	}
}

func hasAuthHeader(c echo.Context) bool {
	return c.Request().Header.Get(echo.HeaderAuthorization) != ""
}
