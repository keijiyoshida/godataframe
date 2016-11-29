package csv

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/keijiyoshida/godataframe/dataframe"
)

// Read reads CSV data from r, creates data frame data and returns it.
func Read(r io.Reader, config dataframe.Config) (*dataframe.DataFrame, error) {
	data, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	return dataframe.New(data, config)
}

// ReadFile reads a CSV data file, creates data frame data and returns it.
func ReadFile(path string, config dataframe.Config) (*dataframe.DataFrame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return Read(f, config)
}
