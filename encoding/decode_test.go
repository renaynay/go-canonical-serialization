package encoding

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_unmarshalUint64(t *testing.T) {
	var tests = []struct {
		in       []byte
		expected uint64
	}{
		{
			in:       []byte{5, 0, 0, 0, 0, 0, 0, 0},
			expected: uint64(5),
		},
		{
			in:       []byte{229, 213, 148, 191, 214, 0, 0, 0},
			expected: uint64(922337203685),
		},
		{
			in: []byte{187, 0, 0, 0, 0, 0, 0, 0},
			expected: uint64(187),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out, err := unmarshalUint64(tt.in)
			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of unmarshalUint64 does not match expected value, unsuccessful")
			}
		})
	}
}

func Test_unmarshalUint32(t *testing.T) {
	var tests = []struct {
		in       []byte
		expected uint32
	}{
		{
			in:       []byte{5, 0, 0, 0},
			expected: uint32(5),
		},
		{
			in:       []byte{15, 192, 249, 54},
			expected: uint32(922337295),
		},
		{
			in: []byte{187, 0, 0, 0},
			expected: uint32(187),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out, err := unmarshalUint32(tt.in)
			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of unmarshalUint64 does not match expected value, unsuccessful")
			}
		})
	}
}



func Test_unmarshalItemList(t *testing.T) { //TODO: turn this into a table test
	expected := [][]byte{
		[]byte{1, 2, 3, 5, 6, 7},
		[]byte{54, 32, 7},
	}

	marshaled := marshalItemList(expected)
	unmarshaled := unmarshalItemList(marshaled)

	if !reflect.DeepEqual(expected, unmarshaled) {
		t.Error("return value of unmarshalItemList does not match expected value, unsuccessful")
	}
}
