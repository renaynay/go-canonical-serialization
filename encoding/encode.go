package encoding

import (
	"bytes"
	"encoding/binary"
)

func EncodeU64(in uint64, out []byte) {
	binary.LittleEndian.PutUint64(out, in)
}

func EncodeU32(in uint32, out []byte) {
	binary.LittleEndian.PutUint32(out, in)
}

func EncodeU8(in uint8) byte {
	return byte(in)
}

func EncodeItem(in []byte) []byte {
	length := make([]byte, 4)

	binary.LittleEndian.PutUint32(length, uint32(len(in)))

	return append(length, in...)
}

func EncodeItemList(in [][]byte) []byte {
	length := make([]byte, 4)

	binary.LittleEndian.PutUint32(length, uint32(len(in)))

	b := bytes.Join(in, nil)

	return append(length, b...)
}