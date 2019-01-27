package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var input string
var output string
var mode bool
var data_len int
var total_len int

func init() {
	var create bool
	var extract bool
	flag.BoolVar(&create, "c", true, "Input mode")
	flag.BoolVar(&extract, "x", false, "Extract mode")
	flag.StringVar(&output, "f", "", "Output file")
	flag.Parse()
	input = flag.Arg(0)
	if input == "" {
		log.Println(flag.Args())
		return
	}
	if create && !extract {
		mode = true
	} else if extract && !create {
		mode = false
	}
	if output == "" && mode {
		output = input + ".huf"
	} else if output == "" && !mode {
		output = input + ".out"
	}
}

func main() {
	if mode {
		// if create
		file, err := ioutil.ReadFile(input)
		if err != nil {
			log.Printf("input: %s\toutput: %s\n", input, output)
			log.Fatal("Read File Error")
			return
		}
		data := NewData(file) //開始編碼
		data.Save(output)
		lf := (len(file))
		log.Printf("input: %d\toutput: %d\tTree: %d\tratio: %.2f%%\n", lf, total_len, total_len-data_len, (100 * float64(total_len) / float64(lf)))
	} else {
		data := new(Data)
		data.Read(input)
		var file []byte = LoadData(data)
		ioutil.WriteFile(output, file, 0644)
	}
}
