package encoding

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_marshalUint64(t *testing.T) {
	var tests = []struct {
		in       uint64
		expected []byte
	}{
		{in: 957386, expected: []byte{202, 155, 14, 0, 0, 0, 0, 0}},
		{in: 48, expected: []byte{48, 0, 0, 0, 0, 0, 0, 0}},
		{in: 3, expected: []byte{3, 0, 0, 0, 0, 0, 0, 0}},
		{in: 2859284759284735, expected: []byte{255, 139, 99, 28, 129, 40, 10, 0}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := make([]byte, 8)

			err := marshalUint64(tt.in, out)
			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of marshalUint64 does not match expected value")
			}
		})
	}
}

func Benchmark_marshalUint64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := make([]byte, 8)
		marshalUint64(uint64(n), b)
	}
}

func Test_marshalUint32(t *testing.T) {
	var tests = []struct {
		in       uint32
		expected []byte
	}{
		{in: 957386, expected: []byte{202, 155, 14, 0}},
		{in: 48, expected: []byte{48, 0, 0, 0}},
		{in: 3, expected: []byte{3, 0, 0, 0}},
		{in: 2859234735, expected: []byte{175, 117, 108, 170}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := make([]byte, 4)

			err := marshalUint32(tt.in, out)
			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of marshalUint32 does not match expected value")
			}
		})
	}
}

func Benchmark_marshalUint32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := make([]byte, 4)
		marshalUint32(uint32(n), b)
	}
}

func Test_marshalUint8(t *testing.T) {
	var tests = []struct {
		in       uint8
		expected byte
	}{
		{in: 95, expected: byte(95)},
		{in: 48, expected: byte(48)},
		{in: 3, expected: byte(3)},
		{in: 28, expected: byte(28)},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(marshalUint8(tt.in), tt.expected) {
				t.Error("return value of marshalUint8 does not match expected value")
			}
		})
	}
}

func Test_marshalItem(t *testing.T) {
	var tests = []struct {
		in       []byte
		expected []byte
	}{
		{in: []byte{1, 2, 3, 4, 5, 6, 8}, expected: []byte{7, 0, 0, 0, 1, 2, 3, 4, 5, 6, 8}},
		{in: []byte{45, 123, 43, 82, 183}, expected: []byte{5, 0, 0, 0, 45, 123, 43, 82, 183}},
		{in: []byte{145, 128, 87, 2}, expected: []byte{4, 0, 0, 0, 145, 128, 87, 2}},
		{in: []byte{4, 56}, expected: []byte{2, 0, 0, 0, 4, 56}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(marshalItem(tt.in), tt.expected) {
				t.Error("return value of EncodeByteArray does not match expected value, unsuccessful")
			}
		})
	}
}

func Test_marshalItemLIst(t *testing.T) {
	var tests = []struct {
		in       [][]byte
		expected []byte
	}{
		{in: [][]byte{
			[]byte{1, 2, 3, 4},
			[]byte{45, 123},
			[]byte{24, 1, 83, 9},
			[]byte{2},
		}, expected: []byte{
			4, 0, 0, 0, 4, 0, 0, 0, 1, 2, 3, 4, 2, 0, 0, 0, 45, 123, 4, 0, 0, 0, 24, 1, 83, 9, 1, 0, 0, 0, 2,
		}},
		{in: [][]byte{
			[]byte{1, 2, 3, 4},
			[]byte{45, 123},
		}, expected: []byte{
			2, 0, 0, 0, 4, 0, 0, 0, 1, 2, 3, 4, 2, 0, 0, 0, 45, 123,
		}},
		{in: [][]byte{
			[]byte{24, 1, 83, 9},
		}, expected: []byte{
			1, 0, 0, 0, 4, 0, 0, 0, 24, 1, 83, 9,
		}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := marshalItemList(tt.in)

			if !reflect.DeepEqual(out, tt.expected) {
				t.Error("return value of marshalItemList did not match expected value, unsuccessful")
			}
		})
	}
}
