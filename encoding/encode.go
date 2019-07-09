package encoding

import (
	"encoding/binary"
)

func marshalUint64(in uint64) []byte {
	out := make([]byte, 8)

	binary.LittleEndian.PutUint64(out, in)

	return out
}

func marshalUint32(in uint32) []byte {
	out := make([]byte, 4)

	binary.LittleEndian.PutUint32(out, in)

	return out
}

func marshalUint8(in uint8) byte {
	return byte(in)
}

func marshalItem(in []byte) []byte {
	length := make([]byte, 4)

	binary.LittleEndian.PutUint32(length, uint32(len(in)))

	return append(length, in...)
}

func marshalItemList(in [][]byte) []byte {
	out := make([]byte, 4)

	binary.LittleEndian.PutUint32(out, uint32(len(in)))

	for item := range in {
		marshaledItem := marshalItem(in[item])
		out = append(out, marshaledItem...)
	}

	return out
}

func marshalMap() { // TODO: finish func
}
