package tsv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/eyecuelab/kit/fileurl"
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
	errCannotReadURL     = Error{"fileurl.downloadtemp: cannot read url"}
	errWrongElementCount = Error{"line has different number of elmeents than record has fields"}
)

//FromPaths parses the path(s) to see whether they are URLs or local paths,
//downloads the file(s) if necessary, then parses them and returns the records
func FromPaths(paths ...string) (records []Record, err error) {
	for _, path := range paths {
		r, err := FromPath(path)
		if err != nil {
			return records, err
		}
		records = append(records, r...)
	}
	return records, nil
}

//FromPath parses a path to see whether it is a URL or local path,
//downloads the file if necessary, then parses it and returns the records
func FromPath(path string) (records []Record, err error) {
	file, err := asFile(path)
	if err != nil {
		return nil, errorF("asFile: %v", err)
	}
	defer file.Close()
	return Parse(file)
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

//asFile takes a URL or local path, downloads if necessary, and returns a file handle
func asFile(s string) (file *os.File, err error) {
	if fileurl.IsFileURL(s) {
		file, err = fileurl.DownloadTemp(s, "factual")
		if err != nil {
			return nil, errCannotReadURL
		}
		return file, nil
	}

	file, err = os.Open(s)
	if err != nil {
		return nil, Error{err.Error()}
	}
	return file, nil
}

//Parse an io.Reader and extract the Records.
func Parse(file io.Reader) (records []Record, err error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	labels := strings.Fields(scanner.Text())
	fmt.Println(labels)
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

func (record Record) getFloat64(key string) (float64, error) {
	s, ok := record[key]
	if !ok {
		return 0, errKeyDoesNotExist
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errCannotConvertKey
	}
	return val, nil
}
