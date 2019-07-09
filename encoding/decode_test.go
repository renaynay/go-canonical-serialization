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
			in:       []byte{187, 0, 0, 0, 0, 0, 0, 0},
			expected: uint64(187),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out, err := unmarshalUint64(tt.in)
			if err != nil {
				t.Error(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of unmarshalUint64 does not match expected value, unsuccessful")
			}
		})
	}
}

func Benchmark_unmarshalUint64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarshalUint64([]byte{64, 0, 0, 0, 0, 0, 0, 0})
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
			in:       []byte{187, 0, 0, 0},
			expected: uint32(187),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out, err := unmarshalUint32(tt.in)
			if err != nil {
				t.Error(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of unmarshalUint32 does not match expected value, unsuccessful")
			}
		})
	}
}

func Benchmark_unmarshalUint32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarshalUint32([]byte{64, 0, 0, 0})
	}
}

func Test_unmarshalUint8(t *testing.T) {
	var tests = []struct {
		in       byte
		expected uint8
	}{
		{in: byte(79), expected: uint8(79)},
		{in: byte(7), expected: uint8(7)},
		{in: byte(123), expected: uint8(123)},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(unmarshalUint8(tt.in), tt.expected) {
				t.Error("return value of unmarshalUint8 does not match expected value, unsuccessful")
			}
		})
	}
}

func Benchmark_unmarshalUint8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarshalUint8(byte(3))
	}
}

func Test_unmarshalItem(t *testing.T) {
	var tests = []struct {
		in []byte
	}{
		{
			in: []byte{1, 2, 3, 4},
		},
		{
			in: []byte{12, 34, 123, 45},
		},
		{
			in: []byte{4, 4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			marshaled := marshalItem(tt.in)

			if !reflect.DeepEqual(tt.in, unmarshalItem(marshaled)) {
				t.Error("return value of unmarshalItem does not match expected value, unsuccessful")
			}
		})
	}
}

func Benchmark_unmarshalItem(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarshalItem([]byte{3, 0, 0, 0, 4, 4, 5})
	}
}

func Test_unmarshalItemList(t *testing.T) {
	var tests = []struct {
		expected [][]byte
	}{
		{
			expected: [][]byte{
				[]byte{1, 2, 3, 5, 6, 7},
				[]byte{54, 32, 7},
			},
		},
		{
			expected: [][]byte{
				[]byte{54, 7},
				[]byte{1, 2},
				[]byte{1, 2, 3, 5, 6, 7, 23, 54},
			},
		},
		{
			expected: [][]byte{
				[]byte{1, 2, 3, 5, 6, 7},
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			marshaled := marshalItemList(tt.expected)

			if !reflect.DeepEqual(unmarshalItemList(marshaled), tt.expected) {
				t.Error("return value of unmarshalItemList does not match expected value, unsuccessful")
			}
		})
	}
}

func Benchmark_unmarshalItemList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarshalItemList([]byte{2, 0, 0, 0, 6, 0, 0, 0, 1, 2, 3, 5, 6, 7, 3, 0, 0, 0, 54, 32, 7})
	}
}
