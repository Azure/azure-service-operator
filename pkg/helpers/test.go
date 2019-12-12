package helpers

import (
	"bytes"
	"encoding/gob"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
)

// test helpers
func GenerateGroupName(identifier string) string {
	return "ci-aso-" + config.BuildID() + "-" + identifier
}

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
