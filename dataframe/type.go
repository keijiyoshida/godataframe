package dataframe

// Type represents a type of a data frame element.
type Type int

// Types of a data frame element.
const (
	Float64 Type = iota
	String
)
