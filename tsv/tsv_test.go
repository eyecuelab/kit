package tsv

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

const (
	johnFirst               = `john`
	johnLast                = `doe`
	johnAge                 = `22`
	janeFirst               = `jane`
	janeLast                = `doe`
	janeAge                 = `58`
	last, lng               = 22.3, -22.4
	firstLabel              = "first"
	lastLabel               = "last"
	ageLabel                = "age"
	testFileName            = "test.tsv"
	localURL                = `http://localhost:3000/test.tsv`
	ownerAllUserReadExecute = os.FileMode(0755)
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

func testTSV(args ...string) string {
	return strings.Join(args, "\t")
}

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
	return ioutil.WriteFile("test.tsv", []byte(testTSVBody), ownerAllUserReadExecute)
}

func init() {
	if err := createTestFile(); err != nil {
		log.Fatalf("could not create test file")
	}
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
			args: args{line: testTSV(johnDoe...), labels: testLabels},
			want: Record{"first": johnFirst, "last": johnLast, "age": johnAge},
			ok:   true,
		}, {
			name: "wrong length",
			args: args{line: testTSV(`"hello"`, `"my"`, `"name"`, `"is"`), labels: testLabels},
			want: nil,
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := parseLine(tt.args.line, tt.args.labels)
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

func Test_FromPath(t *testing.T) {
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
			gotRecords, err := FromPath(tt.args.path)
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

type byFirstName []Record

func (r byFirstName) Less(i, j int) bool {
	fI, _ := r[i]["first"]
	fJ, _ := r[j]["first"]
	return fI < fJ
}

func (r byFirstName) Len() int {
	return len(r)
}

func (r byFirstName) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func chanToSlice(ch <-chan Record) []Record {
	var records []Record
	for r := range ch {
		records = append(records, r)
	}
	return records
}

func TestStreamFromPaths(t *testing.T) {
	type args struct {
		out   chan Record
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Record
		wantErr bool
	}{
		{
			name: "ok - one path",
			args: args{paths: []string{testFileName}, out: make(chan Record)},
			want: []Record{johnRecord, janeRecord},
		}, {
			name: "ok - two paths",
			args: args{paths: []string{testFileName, localURL}, out: make(chan Record, 2)},
			want: []Record{johnRecord, johnRecord, janeRecord, janeRecord},
		}, {
			name:    "err - bad filepath",
			args:    args{paths: []string{"somebadfilename"}, out: make(chan Record)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StreamFromPaths(tt.args.out, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("StreamFromPaths() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := chanToSlice(tt.args.out)
			sort.Sort(byFirstName(got))
			sort.Sort(byFirstName(tt.want))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamFromPaths(): \n got %v: \n want %v", got, tt.want)
			}

		})
	}
}

func Test_labels_ParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		l     labels
		args  args
		want  Record
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.l.ParseLine(tt.args.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("labels.ParseLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("labels.ParseLine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []Record
		wantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := FromPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("FromPath() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func TestStreamFromBindataPaths(t *testing.T) {
	type args struct {
		out   chan Record
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StreamFromBindataPaths(tt.args.out, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("StreamFromBindataPaths() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_asBinReadCloser(t *testing.T) {
}

func Test_parseStream(t *testing.T) {
}

func Test_parseStreams(t *testing.T) {
}
