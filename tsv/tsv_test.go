package tsv

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	johnFirst                  = `john`
	johnLast                   = `doe`
	johnAge                    = `22`
	janeFirst                  = `jane`
	janeLast                   = `doe`
	janeAge                    = `58`
	last, lng                  = 22.3, -22.4
	firstLabel                 = "first"
	lastLabel                  = "last"
	ageLabel                   = "age"
	testFileName               = "test.tsv"
	localURL                   = `http://localhost:3000/test.tsv`
	allPermissions os.FileMode = 0777
)

var (
	wd, _        = os.Getwd()
	testFilePath = wd + "/" + testFileName
	testLabels   = []string{firstLabel, lastLabel, ageLabel}
	johnDoe      = []string{johnFirst, johnLast, johnAge}
	janeDoe      = []string{janeFirst, janeLast, janeAge}
	testTSVBody  = strings.Join([]string{
		strings.Join(testLabels, "\t"),
		strings.Join(johnDoe, "\t"),
		strings.Join(janeDoe, "\t"),
	},
		"\n",
	)

	johnRecord = Record{
		firstLabel: johnFirst,
		lastLabel:  johnLast,
		ageLabel:   johnAge,
	}
	janeRecord = Record{
		firstLabel: janeFirst,
		lastLabel:  janeLast,
		ageLabel:   janeAge,
	}
)

func serveFromMockTSVServer() http.Handler {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", fs)
	log.Println("serving...")
	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Println("serving failed")
		}
	}()
	return fs
}

func createTestFile() error {
	return ioutil.WriteFile("test.tsv", []byte(testTSVBody), os.FileMode(0755))
}

func init() {
	if err := createTestFile(); err != nil {
		log.Fatalf("could not create test file")
	}
	time.Sleep(10 * time.Millisecond)
	serveFromMockTSVServer()

}

func TestRecord_Float64(t *testing.T) {
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
			got, err := tt.record.Float64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Record.Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Record.Float64() = %v, want %v", got, tt.want)
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
	wd, _ := os.Getwd()

	tests := []struct {
		name       string
		args       args
		wantString string
		wantErr    bool
	}{
		{
			name:       "good url",
			args:       args{localURL},
			wantString: testTSVBody,
		}, {
			name:    "does not exist",
			args:    args{"thisseemslike.agoodurl.com/right.tsv"},
			wantErr: true,
		},
		{
			name:       "good file",
			wantString: testTSVBody,
			args:       args{wd + "/" + testFileName},
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
			var got bytes.Buffer
			io.Copy(&got, gotReadCloser)
			if !reflect.DeepEqual(got.String(), tt.wantString) {
				t.Errorf("asReadCLoser(): got %v, want %v", got.String(), tt.wantString)
			}

		})

	}
}

func Test_parseFromPath(t *testing.T) {
	want := []Record{johnRecord, janeRecord}
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []Record
		wantErr     bool
	}{
		{
			name:        "ok website",
			args:        args{localURL},
			wantRecords: want,
		}, {
			name:        "ok local file",
			args:        args{testFileName},
			wantRecords: want,
		}, {
			name:    "bad path",
			args:    args{"badfile.asdasd"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := parseFromPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("parseFromPath() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	mismatched := bytes.NewBufferString("first\tlast\tage\njohn\tdoe\t22\njane\tdoe\n")
	tests := []struct {
		name        string
		args        args
		wantRecords []Record
		wantErr     bool
	}{
		{
			name:        "ok",
			args:        args{reader: bytes.NewBufferString(testTSVBody)},
			wantRecords: []Record{johnRecord, janeRecord},
		}, {
			name:    "mistmatched records",
			args:    args{reader: mismatched},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := Parse(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("Parse() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
