package dataframe

import "errors"

// Errors
var (
	ErrInvalidTypesLength = errors.New("the length of types does not match with the one of item names")
	ErrInvalidType        = errors.New("invalid type")
	ErrDuplicatedItemName = errors.New("duplicated itemName")
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

// newItemNames creates a new slice, copies the source slice to it and returns it.
func newItemNames(srcItemNames []string) []string {
	itemNames := make([]string, len(srcItemNames))

	copy(itemNames, srcItemNames)

	return itemNames
}

// newTypes creates a new item name - type map and returns it.
func newTypes(srcItemNames []string, srcTypes []Type) (map[string]Type, error) {
	if len(srcItemNames) != len(srcTypes) {
		return nil, ErrInvalidTypesLength
	}

	types := make(map[string]Type)

	for i, itemName := range srcItemNames {
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
