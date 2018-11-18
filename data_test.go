package main

import (
	"flag"
	"io/ioutil"
	"testing"
)

var path string
var name string

func init() {
	flag.StringVar(&path, "path", "test.txt", "input the file path")
	flag.StringVar(&name, "name", "", "file output name")
	flag.Parse()
}

func TestFile(t *testing.T) {
	file, err := ioutil.ReadFile(path)
	f_len := len(file)
	if err != nil {
		t.Error(err.Error())
	}
	e := NewData(file)
	if name == "" {
		name = path + ".huf"
	}
	err = e.Save(name)
	if err != nil {
		t.Error(err.Error())
	}
	d := new(Data)
	err = d.Read(name)
	if err != nil {
		t.Error(err.Error())
	}
	result := LoadData(d)
	err = ioutil.WriteFile("output/"+path, result, 0644)
	if err != nil {
		t.Error(err.Error())
	}
	r_len := len(result)
	t.Logf("file length: %d\tresult length: %d\n", f_len, r_len)
}
