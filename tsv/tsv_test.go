package tsv

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	johnFirst = `"john"`
	johnLast  = `"doe"`
	johnAge   = `22`
	janeFirst = `"jane"`
	janeLast  = `"doe"`
	janeAge   = `58`
	last, lng = 22.3, -22.4
)

var (
	testLabels = []string{"first", "last", "age"}
	johnDoe    = []string{johnFirst, johnLast, johnAge}
	janeDoe    = []string{johnFirst, johnLast, johnAge}
)

func TestRecord_getFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		record  Record
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:   "ok",
			record: Record{"latitude": "-22.34"},
			args:   args{"latitude"},
			want:   -22.34,
		}, {
			name:    "key dne",
			record:  Record{},
			args:    args{"latitude"},
			wantErr: true,
		}, {
			name:    "key bad float",
			record:  Record{"latitude": "aaddmm"},
			args:    args{"latitude"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.record.getFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Record.getFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Record.getFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func tsv(s ...string) string {
	return strings.Join(s, "\t")
}

func TestParseLine(t *testing.T) {
	type args struct {
		line   string
		labels []string
	}
	tests := []struct {
		name string
		args args
		want Record
		ok   bool
	}{
		{
			name: "ok",
			args: args{line: tsv(johnDoe...), labels: testLabels},
			want: Record{"first": johnFirst, "last": johnLast, "age": johnAge},
			ok:   true,
		}, {
			name: "wrong length",
			args: args{line: tsv(`"hello"`, `"my"`, `"name"`, `"is"`), labels: testLabels},
			want: nil,
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ParseLine(tt.args.line, tt.args.labels)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("ParseLine() ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}

func Test_asFile(t *testing.T) {
	type args struct {
		s string
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
			gotFile, err := asFile(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("asFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("asFile() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
