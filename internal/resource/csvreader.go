package resource

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Implementation of Rule Reader for reading game rules from csv files.
type CsvRuleReader struct {
	Comma, Comment rune
}

var cvsRuleReaderInstance *CsvRuleReader

// Singelton; not thread-safe.
func GetCsvRuleReader(comma, comment rune) RuleReader {
	if cvsRuleReaderInstance == nil {
		cvsRuleReaderInstance = &CsvRuleReader{comma, comment}
	}
	return cvsRuleReaderInstance
}

// ReadAll reads records from the csv file by url.
// Csv format: question; answer.
func (csvReader *CsvRuleReader) ReadAll(filepath string) (records [][]string, err error) {
	f, e := os.Open(filepath)
	if e != nil {
		return [][]string{}, fmt.Errorf("csv file cannot be opened :: %s", filepath)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = csvReader.Comma
	r.Comment = csvReader.Comment

	data, e := r.ReadAll()
	if e != nil {
		return [][]string{}, fmt.Errorf("csv file [%s] cannot be parsed :: %s", filepath, e)
	}
	return data, nil
}
