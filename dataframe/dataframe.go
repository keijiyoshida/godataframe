package dataframe

import (
	"bytes"
	"errors"
	"strconv"
)

// Errors
var (
	ErrInvalidTypesLen       = errors.New("the length of types does not match with the one of item names")
	ErrInvalidType           = errors.New("invalid type")
	ErrDuplicatedItemName    = errors.New("duplicated itemName")
	ErrNoData                = errors.New("no data")
	ErrInvalidDataColsNum    = errors.New("invalid number of data columns")
	ErrItemNameAlreadyExists = errors.New("itemName already exists")
	ErrItemNameNotExist      = errors.New("itemName does not exist")
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
	return &DataFrame{df.bd, df.fromRowIdx, min(df.fromRowIdx+n, df.toRowIdx)}
}

// Tail creates a new data frame which has last n rows of
// the original data frame.
func (df *DataFrame) Tail(n int) *DataFrame {
	return &DataFrame{df.bd, max(df.toRowIdx-n, df.fromRowIdx), df.toRowIdx}
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

// AppendFloat64ColFromStringCol creates a float64 column from a string column and
// appends it to the data frame.
func (df *DataFrame) AppendFloat64ColFromStringCol(itemName, srcItemName string, convert func(string) (float64, error)) error {
	if _, exist := df.bd.stringCols[itemName]; exist {
		return ErrItemNameAlreadyExists
	}

	if _, exist := df.bd.float64Cols[itemName]; exist {
		return ErrItemNameAlreadyExists
	}

	stringCol, exist := df.bd.stringCols[srcItemName]
	if !exist {
		return ErrItemNameNotExist
	}

	n := len(stringCol)

	float64Col := make([]float64, n)

	ch := make(chan error, numConcurrency)

	d := divUp(n, numConcurrency)

	for i := 0; i < numConcurrency; i++ {
		from := d * i
		to := min(d*(i+1), n)

		go setFloat64FromString(float64Col, stringCol, from, to, convert, ch)
	}

	errs := make([]error, 0, numConcurrency)

	for i := 0; i < numConcurrency; i++ {
		err := <-ch
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return &MultiError{",", errs}
	}

	df.bd.itemNames = append(df.bd.itemNames, itemName)
	df.bd.types[itemName] = Float64
	df.bd.float64Cols[itemName] = float64Col

	return nil
}

// setFloat64FromString creates a float64 data from a string data and
// appends it to the slice.
func setFloat64FromString(float64Col []float64, stringCol []string, from, to int, convert func(string) (float64, error), ch chan<- error) {
	for i := from; i < to; i++ {
		f, err := convert(stringCol[i])
		if err != nil {
			ch <- err
			return
		}

		float64Col[i] = f
	}

	ch <- nil
}

// New creates and returns a data frame.
func New(data [][]string, config Config) (*DataFrame, error) {
	bd, err := newBaseData(data, config)
	if err != nil {
		return nil, err
	}

	return &DataFrame{bd, 0, bd.rowNum()}, nil
}
