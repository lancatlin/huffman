package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

type Data struct {
	Tree map[byte]uint64
	Data []byte
}

func (d Data) Save(path string) error {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(d)
	if err != nil {
		return err
	}
	total_len = len(network.Bytes())
	err = ioutil.WriteFile(path, network.Bytes(), 0644)
	return err
}

func (d *Data) Read(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	network := bytes.NewBuffer(file)
	dec := gob.NewDecoder(network)
	err = dec.Decode(d)
	if err != nil {
		return err
	}
	return nil
}

func NewData(data []byte) *Data {
	d := new(Data)
	tree := Huffman(Count(data))
	d.Data = Encode(tree, data)
	d.Tree = Treeto64(tree)
	data_len = len(d.Data)
	return d
}

func LoadData(data *Data) []byte {
	return Decode(UintToTree(data.Tree), data.Data)
}

func Treeto64(tree map[byte][]byte) (output map[byte]uint64) {
	output = make(map[byte]uint64)
	for key, value := range tree {
		output[key] = SliceTo64(value)
	}
	return
}

func UintToTree(tree map[byte]uint64) (output map[byte][]byte) {
	output = make(map[byte][]byte)
	for key, value := range tree {
		output[key] = UintToSlice(value)
	}
	return
}
