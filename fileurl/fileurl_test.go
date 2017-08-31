package fileurl

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestIsFileURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileURL(tt.args.path); got != tt.want {
				t.Errorf("IsFileURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantDst string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := &bytes.Buffer{}
			gotN, err := Copy(tt.args.url, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Copy() = %v, want %v", gotN, tt.wantN)
			}
			if gotDst := dst.String(); gotDst != tt.wantDst {
				t.Errorf("Copy() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestDownloadTemp(t *testing.T) {
	type args struct {
		url    string
		prefix string
	}
	tests := []struct {
		name     string
		args     args
		wantFile *os.File
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, err := DownloadTemp(tt.args.url, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadTemp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("DownloadTemp() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
