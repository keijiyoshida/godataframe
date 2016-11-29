package dataframe

import "errors"

// Errors
var (
	ErrInvalidTypesLength = errors.New("the length of types does not match with the one of item names")
	ErrInvalidType        = errors.New("invalid type")
	ErrDuplicatedItemName = errors.New("duplicated itemName")
	ErrNoItemNames        = errors.New("no item names")
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
	itemNames, err := newItemNames(data, config)
	if err != nil {
		return nil, err
	}

	types, err := newTypes(config.ItemNames, config.Types)
	if err != nil {
		return nil, err
	}

	return &DataFrame{itemNames, types, nil, nil}, nil
}

// getSrcItemNames extracts source item names and returns them.
func getSrcItemNames(data [][]string, config Config) ([]string, error) {
	var srcItemNames []string

	if config.UseFirstRowAsItemNames {
		if len(data) < 1 {
			return nil, ErrNoItemNames
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
		return nil, ErrInvalidTypesLength
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

// newColumns creates string and float64 columns and returns them.
func newColumns() (map[string][]string, map[string][]float64) {
	return nil, nil
}
