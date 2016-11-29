package dataframe

import (
	"bytes"
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
	bd         *baseData
	fromRowIdx int // inclusive
	toRowIdx   int // exclusive
}

// RowNum returns the number of rows.
func (df *DataFrame) RowNum() int {
	return df.toRowIdx - df.fromRowIdx
}

// ColNum returns the number of columns.
func (df *DataFrame) ColNum() int {
	return len(df.bd.itemNames)
}

// Head creates a new data frame which has top n rows of
// the original data frame.
func (df *DataFrame) Head(n int) *DataFrame {
	return &DataFrame{df.bd, df.fromRowIdx, min(df.bd.rowNum(), df.fromRowIdx+n)}
}

// String returns the string expression of the data frame.
func (df *DataFrame) String() string {
	bf := bytes.NewBufferString("")

	for i, itemName := range df.bd.itemNames {
		if i > 0 {
			bf.WriteRune(' ')
		}

		bf.WriteString(itemName)
	}

	bf.WriteRune('\n')

	for i, n := 0, min(maxPrintRows, (df.toRowIdx-df.fromRowIdx)); i < n; i++ {
		if i > 0 {
			bf.WriteRune('\n')
		}

		for j, itemName := range df.bd.itemNames {
			if j > 0 {
				bf.WriteRune(' ')
			}

			t := df.bd.types[itemName]

			if t == String {
				bf.WriteString(df.bd.stringCols[itemName][i+df.fromRowIdx])
			} else {
				bf.WriteString(strconv.FormatFloat(df.bd.float64Cols[itemName][i+df.fromRowIdx], 'f', 8, 64))
			}
		}
	}

	return bf.String()
}

// New creates and returns a data frame.
func New(data [][]string, config Config) (*DataFrame, error) {
	bd, err := newBaseData(data, config)
	if err != nil {
		return nil, err
	}

	return &DataFrame{bd, 0, bd.rowNum()}, nil
}
