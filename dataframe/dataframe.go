package dataframe

// DataFrame represents a data frame.
type DataFrame struct {
	header *header
	body   *body
}

// New creates and returns a data frame.
func New(data [][]string, config Config) (*DataFrame, error) {
	return nil, nil
}
