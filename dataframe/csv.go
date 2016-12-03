package dataframe

import (
	"encoding/csv"
	"io"
	"os"
)

// ReadCSV reads CSV data from r, creates data frame data and returns it.
func ReadCSV(r io.Reader, config Config) (*DataFrame, error) {
	data, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	return New(data, config)
}

// ReadCSVFile reads a CSV data file, creates data frame data and returns it.
func ReadCSVFile(path string, config Config) (*DataFrame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ReadCSV(f, config)
}
