package middleware

import (
	"github.com/labstack/echo"
	emw "github.com/labstack/echo/middleware"
)

type (
	// ApiErrorConfig defines the config for ApiError middleware.
	TemplateConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper emw.Skipper
	}
)

var (
	// DefaultApiErrorConfig is the default ApiError middleware config.
	DefaultTemplateConfig = TemplateConfig{
		Skipper: emw.DefaultSkipper,
	}
)

// ApiError returns a middleware that logs converts errors to json api errors.
func ApiError() echo.MiddlewareFunc {
	return TemplateWithConfig(DefaultTemplateConfig)
}

func TemplateWithConfig(config TemplateConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultTemplateConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			return next(c)
		}
	}
}
