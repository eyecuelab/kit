package tsv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Record represents a single line of a TSV
type Record map[string]string

type Error struct {
	msg string
}

func (err Error) Error() string {
	return err.msg
}

func errorF(format string, a ...interface{}) Error {
	return Error{fmt.Sprintf(format, a...)}
}

var (
	errKeyDoesNotExist   = Error{"key does not exist"}
	errCannotConvertKey  = Error{"cannot convert key"}
	errWrongElementCount = Error{"line has different number of elmeents than record has fields"}
)

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

//ParseLine parses a single line of a TSV using the given labels
func ParseLine(line string, labels []string) (Record, bool) {
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
	labels := strings.Fields(scanner.Text())
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		record, ok := ParseLine(scanner.Text(), labels)
		if !ok {
			return nil, errWrongElementCount
		}
		records = append(records, record)
	}
	return records, nil
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
