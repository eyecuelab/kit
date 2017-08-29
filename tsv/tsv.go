package tsv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/eyecuelab/kit/fileurl"
	"github.com/eyecuelab/kit/tickertape"
)

//Record represents a single line of a TSV
type Record map[string]string

//FromPath parses a path to see whether it is a URL or local path,
//downloads the file if necessary, then parses it and returns the records
func FromPath(paths ...string) (records []Record, err error) {
	for _, path := range paths {
		file, err := asFile(path)
		if err != nil {
			return nil, fmt.Errorf("asFile: %v", err)
		}
		defer file.Close()
		r, err := Parse(file)
		if err != nil {
			return records, fmt.Errorf("getRecords: parsing file at path %s, %v", path, err)
		}
		records = append(records, r...)
	}
	return records, nil
}

//ParseLine parses a single line of a TSV using the given labels
func ParseLine(line string, labels []string) (Record, bool) {
	split := strings.Split(line, "\t")
	if len(split) != len(labels) {
		fmt.Println(split)
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
		tickertape.Printf("downloading %s from %s", path.Base(s)+path.Ext(s), path.Dir(s))
		file, err = fileurl.DownloadTemp(s, "factual")
		if err != nil {
			return nil, fmt.Errorf("could not read URL: fileurl.DownloadTemp: %v", err)
		}
		return file, nil
	}
	//otherwise, assume is path to local file
	tickertape.Printf("Opening local file %s in %s", path.Base(s)+path.Ext(s), path.Dir(s))
	file, err = os.Open(s)
	if err != nil {
		return nil, fmt.Errorf("could not read local file: %v", err)
	}
	return file, nil
}

//Parse an io.Reader and extract the Records.
func Parse(file io.Reader) (records []Record, err error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	labels := strings.Fields(scanner.Text())
	for scanner.Scan() {
		line := scanner.Text()
		record, ok := ParseLine(scanner.Text(), labels)
		if !ok {
			return nil, fmt.Errorf("could not create record: %d labels, but %d elements in line", len(labels), len(line))
		}
		records = append(records, record)
	}
	return records, nil
}
