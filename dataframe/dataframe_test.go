package dataframe

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newItemNames(t *testing.T) {
	testCases := [][]string{
		{},
		{"0"},
		{"0", "1"},
		{"0", "1", "2"},
		{"0", "1", "2", "3"},
	}

	for _, tc := range testCases {
		got := newItemNames(tc)

		if !reflect.DeepEqual(got, tc) {
			t.Errorf("newItemNames(%v) => %v; want %v", tc, got, tc)
		}

		if fmt.Sprintf("%p", got) == fmt.Sprintf("%p", tc) && len(got) > 0 {
			t.Errorf("newItemNames should return a new slice")
		}
	}
}

func Test_newTypes(t *testing.T) {
	testCases := []struct {
		srcItemNames []string
		srcTypes     []Type
		wantTypes    map[string]Type
		wantErr      error
	}{
		{
			[]string{"0"},
			[]Type{String, Float64},
			nil,
			ErrInvalidTypesLength,
		},
		{
			[]string{"0"},
			[]Type{Type(-1)},
			nil,
			ErrInvalidType,
		},
		{
			[]string{"0", "0"},
			[]Type{String, Float64},
			nil,
			ErrDuplicatedItemName,
		},
		{
			[]string{"0", "1", "2", "3"},
			[]Type{String, Float64, String, Float64},
			map[string]Type{"0": String, "1": Float64, "2": String, "3": Float64},
			nil,
		},
	}

	for _, tc := range testCases {
		gotTypes, gotErr := newTypes(tc.srcItemNames, tc.srcTypes)

		if !reflect.DeepEqual(gotTypes, tc.wantTypes) || gotErr != tc.wantErr {
			t.Errorf("newTypes(%v, %v) => %v, %#v; want %v, %#v",
				tc.srcItemNames, tc.srcTypes, gotTypes, gotErr, tc.wantTypes, tc.wantErr)
		}
	}
}
