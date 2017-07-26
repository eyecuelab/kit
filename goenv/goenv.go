package goenv

import (
	"os"
)

var Env string
var Dev, Test, Prod bool

func init() {
	Env = os.Getenv("GO_ENV")
	if Env == "" {
		Env = "development"
	}

	environs := map[string]*bool{"development": &Dev, "test": &Test, "production": &Prod}
	*environs[Env] = true
}
