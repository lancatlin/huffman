package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var input string
var output string
var create bool
var extract bool

func init() {
	flag.BoolVar(&create, "c", true, "Input mode")
	flag.BoolVar(&extract, "x", false, "Extract mode")
	flag.StringVar(&output, "f", "", "Output file")
	flag.Parse()
	input = flag.Arg(0)
	if input == "" {
		log.Println(flag.Args())
		return
	}
	if output == "" && create {
		output = input + ".huf"
	}
}

func main() {
	if create && !extract {
		file, err := ioutil.ReadFile(input)
		if err != nil {
			log.Printf("input: %s\toutput: %s\n", input, output)
			log.Fatal("Read File Error")
			return
		}
		data := NewData(file) //開始編碼
		data.Save(output)
		lf, lr := (len(file)), (len(data.Data))
		log.Printf("input: %d\toutput: %dratio: %.2f%%\n", lf, lr, (100 * float64(lr) / float64(lf)))
	} else if extract || !create {
		data := new(Data)
		data.Read(input)
		var file []byte = LoadData(data)
		ioutil.WriteFile(output, file, 0644)
	}
}
