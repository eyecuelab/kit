package assets

import "errors"

type (
	AssetGet     func(string) ([]byte, error)
	AssetDir     func(string) ([]string, error)
	AssetManager struct {
		Get AssetGet
		Dir AssetDir
	}
)

var Manager *AssetManager

// Get retrieve static asset from client project data directory. This allows code in kit to use client project data dir
func Get(name string) ([]byte, error) {
	if Manager == nil {
		return nil, errors.New("Manager is not set")
	}

	return Manager.Get(name)
}

func Dir(name string) ([]string, error) {
	if Manager == nil {
		return nil, errors.New("Manager is not set")
	}

	return Manager.Dir(name)
}
