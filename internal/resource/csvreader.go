package resource

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Implementation of Rule Reader for reading game rules from csv files.
// CsvRuleReader is a singleton.
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
func (csvReader *CsvRuleReader) ReadAll(url string) (records [][]string, err error) {
	f, e := os.Open(url)
	if e != nil {
		return [][]string{}, fmt.Errorf("csv file cannot be opened :: %s", url)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = csvReader.Comma
	r.Comment = csvReader.Comment

	data, e := r.ReadAll()
	if e != nil {
		return [][]string{}, fmt.Errorf("csv file [%s] cannot be parsed :: %s", url, e)
	}
	return data, nil
}
