package encoding

import (
	"encoding/binary"
	"errors"
)

func marshalUint64(in uint64, out []byte) error {
	if len(out) != 8 {
		return errors.New("marshalUint64 requires array that is 8 bytes in length")
	}

	binary.LittleEndian.PutUint64(out, in)

	return nil
}

func marshalUint32(in uint32, out []byte) error {
	if len(out) != 4 {
		return errors.New("marshalUint32 requires array that is 4 bytes in length")
	}

	binary.LittleEndian.PutUint32(out, in)

	return nil
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

	for item, _ := range in {
		marshaledItem := marshalItem(in[item])
		out = append(out, marshaledItem...)
	}

	return out
}

func marshalMap() { // TODO: finish func
}
