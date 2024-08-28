package id

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
)

type ID string

func NewID() ID {
	b := make([]byte, 32)
	var r io.Reader = rand.Reader

	if _, err := io.ReadFull(r, b); err != nil {
		panic(err) // This shouldn't happen
	}

	id := hex.EncodeToString(b)
	return ID(id)
}

func (id ID) ShortString() string {
	return string(id)[1:12]
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(id))
}
