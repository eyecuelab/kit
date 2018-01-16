package assets

import (
	"testing"

	"github.com/eyecuelab/kit/copyslice"
	"github.com/stretchr/testify/assert"

	"github.com/eyecuelab/kit/errorlib"
	"github.com/eyecuelab/kit/maputil/stringmap"
)

var testAssets = map[string][]byte{}

const errMissingAsset errorlib.ErrorString = "missing asset"

func mockGet(m map[string]string) AssetGet {
	m = stringmap.Copy(m)
	return func(s string) ([]byte, error) {
		if v, ok := m[s]; ok {
			return []byte(v), nil
		}
		return nil, errMissingAsset
	}
}

func mockDir(m map[string][]string) AssetDir {
	copy := make(map[string][]string)
	for k, v := range m {
		copy[k] = copyslice.String(v)
	}
	return func(s string) ([]string, error) {
		if v, ok := m[s]; ok {
			return v, nil
		}
		return nil, errMissingAsset
	}
}

func revertManager() { Manager = nil }

func TestGet(t *testing.T) {
	const key, want = "foo", "bar"
	defer revertManager()
	_, err := Get(key)
	assert.EqualError(t, err, string(ErrManagerNotSet))

	Manager = new(AssetManager)
	_, err = Get(key)
	assert.EqualError(t, err, string(ErrNoGetFunc))

	get := mockGet(map[string]string{key: want})
	Manager = &AssetManager{Get: get}
	got, err := Get(key)
	assert.NoError(t, err)
	assert.Equal(t, []byte(want), got)

}
func TestDir(t *testing.T) {
	const key = "foo"
	want := []string{"bar", "baz"}
	defer revertManager()
	_, err := Dir(key)
	assert.EqualError(t, err, string(ErrManagerNotSet))

	Manager = new(AssetManager)
	_, err = Dir(key)
	assert.EqualError(t, err, string(ErrNoDirFunc))

	dir := mockDir(map[string][]string{key: want})
	Manager = &AssetManager{Dir: dir}
	got, err := Dir(key)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

}
