package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/eyecuelab/kit/log"
)

// Load loading env vars from env file
func Load(env string, srcPath string, inject ...string) error {
	gopath := os.Getenv("GOPATH")
	path := fmt.Sprintf("%s/src/%s/.env", gopath, srcPath)
	if env != "" {
		path = fmt.Sprintf("%s.%s", path, env)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Print("No .env file found, falling back to actual env vars")
		return nil
	}

	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("Error loading env file")
	}

	return nil
}
