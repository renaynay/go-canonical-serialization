package encoding

import (
	"encoding/binary"
	"errors"
)

func unmarshalUint64(in []byte) (uint64, error) {
	if len(in) != 8 {
		return uint64(0), errors.New("unmarshalUint64 requires array that is 8 bytes in length")
	}

	return binary.LittleEndian.Uint64(in), nil
}

func unmarshalUint32(in []byte) (uint32, error) {
	if len(in) != 4 {
		return uint32(0), errors.New("unmarshalUint32 requires array that is 4 bytes in length")
	}

	return binary.LittleEndian.Uint32(in), nil
}

func unmarshalUint8(in byte) uint8 {
	return uint8(in)
}

func unmarshalItem(in []byte) []byte {
	length := binary.LittleEndian.Uint32(in[:4])

	return in[4:length]
}

func unmarshalItemList(in []byte) [][]byte {
	index := 4

	amount := int(binary.LittleEndian.Uint32(in[:index]))

	var item []byte
	var itemLen int
	var store [][]byte

	for i := 0; i < amount; i++ {
		itemLen = int(binary.LittleEndian.Uint32(in[index : index+4]))

		item = in[index+4 : index+4+itemLen]

		store = append(store, item)

		index = index + 4 + itemLen
	}

	return store
}
