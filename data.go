package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

type Data struct {
	Tree map[byte][]byte
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
	d.Tree = Huffman(Count(data))
	d.Data = Encode(d.Tree, data)
	data_len = len(d.Data)
	return d
}

func LoadData(data *Data) []byte {
	return Decode(data.Tree, data.Data)
}
