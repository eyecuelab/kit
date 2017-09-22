package tsv

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

const (
	johnFirst   = `"john"`
	johnLast    = `"doe"`
	johnAge     = `22`
	janeFirst   = `"jane"`
	janeLast    = `"doe"`
	janeAge     = `58`
	last, lng   = 22.3, -22.4
	firstLabel  = "first"
	lastLabel   = "last"
	ageLabel    = "age"
	mockHTTPURL = "https://somedomain.co.uk/foo/bar.tsv"
)

var (
	testLabels  = []string{firstLabel, lastLabel, ageLabel}
	johnDoe     = []string{johnFirst, johnLast, johnAge}
	janeDoe     = []string{johnFirst, johnLast, johnAge}
	testTSVBody = strings.Join([]string{
		strings.Join(testLabels, "\t"),
		strings.Join(johnDoe, "\t"),
		strings.Join(janeDoe, "\t"),
	},
		"\n",
	)
	testBodyBytes = []byte(testTSVBody)
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
			recover()
		})
	}
}

func Test_asReadCloser(t *testing.T) {
	type args struct {
		s string
	}

	tempFile, err := ioutil.TempFile("", "test")
	if err == nil {
		defer tempFile.Close()
		tempFile.Write(testBodyBytes)
	}

	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", mockHTTPURL, mockHTTPStringResponder)

	tests := []struct {
		name      string
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		{
			name:      "good url",
			args:      args{mockHTTPURL},
			wantBytes: testBodyBytes,
		}, {
			name:    "bad url",
			args:    args{"thisseemslike.agoodurl.com/right.tsv"},
			wantErr: true,
		},
		{
			name: "good file",
			args: args{tempFile.Name()},
		},
		{
			name:    "bad file",
			args:    args{"thisfiledoesnotexist.tsv"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReadCloser, err := asReadCloser(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("asReadCloser() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				return
			}
			var gotBytes []byte
			gotReadCloser.Read(gotBytes)
			fmt.Print(gotBytes)
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("asReadCLoser(): got %v, want %v", gotBytes, tt.wantBytes)
			}

		})

	}
}
