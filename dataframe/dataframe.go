package dataframe

import (
	"errors"
	"strconv"
)

// Errors
var (
	ErrInvalidTypesLen    = errors.New("the length of types does not match with the one of item names")
	ErrInvalidType        = errors.New("invalid type")
	ErrDuplicatedItemName = errors.New("duplicated itemName")
	ErrNoData             = errors.New("no data")
	ErrInvalidDataColsNum = errors.New("invalid number of data columns")
)

// DataFrame represents a data frame.
type DataFrame struct {
	itemNames   []string
	types       map[string]Type
	stringCols  map[string][]string
	float64Cols map[string][]float64
}

// New creates and returns a data frame.
func New(data [][]string, config Config) (*DataFrame, error) {
	itemNames, err := newItemNames(data, config)
	if err != nil {
		return nil, err
	}

	types, err := newTypes(config.ItemNames, config.Types)
	if err != nil {
		return nil, err
	}

	stringCols, float64Cols, err := newCols(itemNames, config, data)
	if err != nil {
		return nil, err
	}

	return &DataFrame{itemNames, types, stringCols, float64Cols}, nil
}

// getSrcItemNames extracts source item names and returns them.
func getSrcItemNames(data [][]string, config Config) ([]string, error) {
	var srcItemNames []string

	if config.UseFirstRowAsItemNames {
		if len(data) < 1 {
			return nil, ErrNoData
		}

		srcItemNames = data[0]
	} else {
		srcItemNames = config.ItemNames
	}

	return srcItemNames, nil
}

// newItemNames creates a new slice, copies the source slice to it and returns it.
func newItemNames(data [][]string, config Config) ([]string, error) {
	srcItemNames, err := getSrcItemNames(data, config)
	if err != nil {
		return nil, err
	}

	itemNames := make([]string, len(srcItemNames))

	copy(itemNames, srcItemNames)

	return itemNames, nil
}

// newTypes creates a new item name - type map and returns it.
func newTypes(itemNames []string, srcTypes []Type) (map[string]Type, error) {
	if len(itemNames) != len(srcTypes) {
		return nil, ErrInvalidTypesLen
	}

	types := make(map[string]Type)

	for i, itemName := range itemNames {
		t := srcTypes[i]

		if !t.valid() {
			return nil, ErrInvalidType
		}

		if _, exist := types[itemName]; exist {
			return nil, ErrDuplicatedItemName
		}

		types[itemName] = t
	}

	return types, nil
}

// newCols creates string and float64 columns and returns them.
func newCols(itemNames []string, config Config, data [][]string) (map[string][]string, map[string][]float64, error) {
	if len(data) < 1 {
		return nil, nil, ErrNoData
	}

	if len(data[0]) != len(itemNames) {
		return nil, nil, ErrInvalidDataColsNum
	}

	if config.UseFirstRowAsItemNames {
		data = data[1:]
	}

	recNum := len(data)

	stringCols := make(map[string][]string)
	float64Cols := make(map[string][]float64)

	for colIdx, itemName := range itemNames {
		switch config.Types[colIdx] {
		case String:
			stringCols[itemName] = newStringCol(colIdx, recNum, data)
		case Float64:
			float64Col, err := newFloat64Col(colIdx, recNum, data)
			if err != nil {
				return nil, nil, err
			}

			float64Cols[itemName] = float64Col
		}
	}

	return stringCols, float64Cols, nil
}

// newStringCol creates and returns string column data.
func newStringCol(colIdx int, recNum int, data [][]string) []string {
	stringCol := make([]string, recNum)

	for i := 0; i < recNum; i++ {
		stringCol[i] = data[i][colIdx]
	}

	return stringCol
}

// newFloatCol creates and returns float64 column data.
func newFloat64Col(colIdx int, recNum int, data [][]string) ([]float64, error) {
	float64Col := make([]float64, recNum)

	for i := 0; i < recNum; i++ {
		f, err := strconv.ParseFloat(data[i][colIdx], 64)
		if err != nil {
			return nil, err
		}

		float64Col[i] = f
	}

	return float64Col, nil
}
