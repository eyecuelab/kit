package assets

import (
	"github.com/eyecuelab/kit/errorlib"
)

type (
	AssetGet     func(string) ([]byte, error)
	AssetDir     func(string) ([]string, error)
	AssetManager struct {
		Get AssetGet
		Dir AssetDir
	}
)

const (
	ErrManagerNotSet errorlib.ErrorString = "asset manager is not set"
	ErrNoGetFunc     errorlib.ErrorString = "asset manager has no get func"
	ErrNoDirFunc     errorlib.ErrorString = "asset manager has no dir func"
)

var Manager *AssetManager

// Get retrieve static asset from client project data directory. This allows code in kit to use client project data dir
func Get(name string) ([]byte, error) {
	if Manager == nil {
		return nil, ErrManagerNotSet
	} else if Manager.Get == nil {
		return nil, ErrNoGetFunc
	}

	return Manager.Get(name)
}

func Dir(name string) ([]string, error) {
	if Manager == nil {
		return nil, ErrManagerNotSet
	} else if Manager.Dir == nil {
		return nil, ErrNoDirFunc
	}

	return Manager.Dir(name)
}
