package goenv

import (
	"os"
)

var Env string
var Dev, Stage, Prod bool

func init() {
	Env = os.Getenv("GO_ENV")
	if Env == "" {
		Env = "development"
	}

	environs := map[string]*bool{"development": &Dev, "staging": &Stage, "production": &Prod}
	*environs[Env] = true
}
