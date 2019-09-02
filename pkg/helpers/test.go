package helpers

import (
	"bytes"
	"encoding/gob"
)

func ToByteArray(x interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	// Stand-in for a buffer connection
	enc := gob.NewEncoder(&buffer)
	// Will write to buffer.
	err := enc.Encode(x)
	if err != nil {
		return []byte{}, err
	}
	return buffer.Bytes(), err
}

func FromByteArray(r []byte, x interface{}) error {
	var buffer bytes.Buffer
	// Stand-in for a buffer connection
	dec := gob.NewDecoder(&buffer)
	// Will write to buffer.
	buffer.Write(r)
	return dec.Decode(x)
}
