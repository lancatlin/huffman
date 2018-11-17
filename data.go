package main

import (
	"bytes"
	"encoding/gob"
)

type Data struct {
	Tree map[byte][]byte
	Last byte
	Data []byte
}

func (d Data) Save(path string) (err error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err = enc.Encode(d)
	if err != nil {
		return
	}
	return
}
