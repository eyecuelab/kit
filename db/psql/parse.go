package psql

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/eyecuelab/kit/assets"
)

type QueryParser struct {
	fileData  []byte
	variables map[string]interface{}
}

// Get is a shortcut for r.Parse(), passing nil as data.
func (qp QueryParser) Get(name string) (string, error) {
	return qp.Parse(name, nil)
}

func (qp QueryParser) Parse(name string, data interface{}) (string, error) {
	t := template.Must(template.New("").Parse(string(qp.fileData)))

	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, name, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (qp QueryParser) ParseWithValueHolders(name string, numfields int, numEntities int) (string, error) {
	return qp.Parse(name, map[string]interface{}{"ValueHolders": namedPqParams(numfields, numEntities)})
}

func (qp *QueryParser) SetFileFromBindata(path string) error {
	fileData, err := assets.Get(path)
	if err != nil {
		return err
	}

	qp.fileData = fileData
	return nil
}

//Builds named params in the format Postgres wants them:  ($1, $2, $3), ($4, $5, $6)...
func namedPqParams(numFields int, numEntities int) string {
	paramGroups := make([]string, 0, numEntities)
	params := make([]string, 0, numFields)
	numTotalParams := numEntities * numFields

	for i := 0; i < numTotalParams; i++ {
		params = append(params, fmt.Sprintf("$%v", i+1))

		//when we reach the number of fields, group the params
		if (i+1)%numFields == 0 {
			paramGroup := fmt.Sprintf("(%s)", strings.Join(params, ","))
			paramGroups = append(paramGroups, paramGroup)
			params = make([]string, 0, numFields)
		}
	}
	return strings.Join(paramGroups, ",")
}
