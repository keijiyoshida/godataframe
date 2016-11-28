package dataframe

// Config represents configuration of a data frame.
type Config struct {
	// ItemNames represents item names of data frame columns.
	ItemNames []string
	// Types represents types of data frame columns.
	Types []Type
	// The first row is used as a header when UseFirstRowAsHeader is set to true.
	UseFirstRowAsHeader bool
}
