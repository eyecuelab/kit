package assets

import "errors"

type AssetGet func(string) ([]byte, error)
type AssetDir func(string) ([]string, error)

// Getter should be set to the Assets func created by bindata
var Getter AssetGet
var Dirrer AssetDir

// Get retrieve static asset from client project data directory. This allows code in kit to use client project data dir
func Get(name string) ([]byte, error) {
	if Getter == nil {
		return nil, errors.New("assets.Getter is not set")
	}

	return Getter(name)
}

func Dir(name string) ([]string, error) {
	if Dirrer == nil {
		return nil, errors.New("assets.Dirrer is not set")
	}

	return Dirrer(name)
}
