package dataframe

import "errors"

// Errors
var (
	ErrUnequalItemNamesTypesLength = errors.New("the length of itemsNames does not equal the length of types")
	ErrInvalidType                 = errors.New("types contains an invalid type")
	ErrDuplicateItemName           = errors.New("itemNames has a duplicate item name")
)

// DataFrame represents a data frame.
type DataFrame struct {
	itemNames      []string
	types          map[string]Type
	stringColumns  map[string][]string
	float64Columns map[string][]float64
}

// New creates and returns a data frame.
func New(data [][]string, config Config) (*DataFrame, error) {
	return nil, nil
}
