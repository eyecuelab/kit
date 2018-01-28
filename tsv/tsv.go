package tsv

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/eyecuelab/kit/assets"
)

//Record represents a single line of a TSV
type (
	Record map[string]string
	labels []string
	Error  string
)

const (
	errKeyDoesNotExist   Error = "key does not exist"
	errCannotConvertKey  Error = "cannot convert key"
	errWrongElementCount Error = "line has different number of elements than record has fields"
)

func (l labels) ParseLine(line string) (Record, bool) { return parseLine(line, l) }
func (err Error) Error() string                       { return string(err) }
func errorF(format string, a ...interface{}) Error    { return Error(fmt.Sprintf(format, a...)) }

//FromPath parses a path to see whether it is a URL or local path,
//downloads the file if necessary, then parses it and returns the records
func FromPath(path string) (records []Record, err error) {
	readCloser, err := asReadCloser(path)
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()
	return Parse(readCloser)
}

func StreamFromPaths(out chan Record, paths ...string) error {
	readClosers := make([]io.ReadCloser, len(paths))
	for i, path := range paths {
		rc, err := asReadCloser(path)
		if err != nil {
			close(out)
			return err
		}
		readClosers[i] = rc
	}
	go parseStreams(out, readClosers...)
	return nil
}

func StreamFromBindataPaths(out chan Record, paths ...string) error {
	readClosers := make([]io.ReadCloser, len(paths))
	for i, path := range paths {
		rc, err := asBinReadCloser(path)
		if err != nil {
			close(out)
			return errorF("StreamFromBinDataPaths: %v", err)
		}
		readClosers[i] = rc
	}
	go parseStreams(out, readClosers...)
	return nil
}

//ParseLine parses a single line of a TSV using the given labels
func parseLine(line string, labels labels) (Record, bool) {
	split := strings.Split(line, "\t")
	if len(split) != len(labels) {
		return nil, false
	}
	record := make(Record)
	for i, label := range labels {
		record[label] = split[i]
	}
	return record, true

}

//asReadCloser bindata path, and returns a ReadCloser containing the information
func asBinReadCloser(s string) (readCloser io.ReadCloser, err error) {
	b, err := assets.Get(s)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewReader(b)), nil
}

//asReadCloser takes a URL or local path, downloads if necessary, and returns a ReadCloser containing the information
func asReadCloser(s string) (readCloser io.ReadCloser, err error) {
	resp, err := http.Get(s)
	if err == nil {
		return resp.Body, nil
	}

	return os.Open(s)
}

//Parse an io.Reader and extract the Records.
func Parse(reader io.Reader) (records []Record, err error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	labels := labels(strings.Fields(scanner.Text()))
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		record, ok := labels.ParseLine(scanner.Text())
		if !ok {
			return nil, errWrongElementCount
		}
		records = append(records, record)
	}
	return records, nil
}

func parseStream(out chan<- Record, readCloser io.ReadCloser) error {
	defer readCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	scanner.Scan()
	labels := labels(strings.Fields(scanner.Text()))
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		record, ok := labels.ParseLine(scanner.Text())
		if !ok {
			return errWrongElementCount
		}
		out <- record
	}
	return nil
}

func parseStreams(out chan<- Record, readClosers ...io.ReadCloser) {
	defer close(out)
	wg := &sync.WaitGroup{}
	for _, rc := range readClosers {
		wg.Add(1)
		go func(r io.ReadCloser) {
			defer wg.Done()
			if err := parseStream(out, r); err != nil {
				log.Println(err)
			}
		}(rc)
	}
	wg.Wait()
}

//Float64 gets the specified key as a Float64, if possible
func (r Record) Float64(key string) (float64, error) {
	s, ok := r[key]
	if !ok {
		return 0, errKeyDoesNotExist
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errCannotConvertKey
	}
	return val, nil
}

//Bool gets the specified key as a Bool, if possible
func (r Record) Bool(key string) (bool, error) {
	s, ok := r[key]
	if !ok {
		return false, errKeyDoesNotExist
	}
	if s == "true" {
		return true, nil
	}
	if s == "false" {
		return false, nil
	}
	return false, errCannotConvertKey
}

//Int gets the specified key as an Int, if possible
func (r Record) Int(key string) (int, error) {
	s, ok := r[key]
	if !ok {
		return 0, errKeyDoesNotExist
	}
	return strconv.Atoi(s)
}

//StringSlice gets the specified key as a []string, if possible
func (r Record) StringSlice(key string) (a []string, err error) {
	s, ok := r[key]
	if !ok {
		return nil, errKeyDoesNotExist
	}
	err = json.Unmarshal([]byte(s), &a)
	return a, err
}

//IntSlice gets the specified key as a  []Int, if possible
func (r Record) IntSlice(key string) (a []int, err error) {
	s, ok := r[key]
	if !ok {
		return nil, errKeyDoesNotExist
	}
	err = json.Unmarshal([]byte(s), &a)
	return a, err
}

//FloatSlice gets the specified key as a []Float64, if possible
func (r Record) FloatSlice(key string) (a []float64, err error) {
	s, ok := r[key]
	if !ok {
		return nil, errKeyDoesNotExist
	}
	err = json.Unmarshal([]byte(s), &a)
	return a, err
}
