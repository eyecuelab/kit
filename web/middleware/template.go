package middleware

import "github.com/labstack/echo"

type (
	// ApiErrorConfig defines the config for ApiError middleware.
	TemplateConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper mw.Skipper
	}
)

var (
	// DefaultApiErrorConfig is the default ApiError middleware config.
	DefaultTemplateConfig = TemplateConfig{
		Skipper: mw.DefaultSkipper,
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

			next(c)
		}
	}
}
