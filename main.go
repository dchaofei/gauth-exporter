package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dchaofei/gauth-exporter/parse"
)

func main() {
	uri := flag.String("uri", "", "uri is otpauth-migration://offline?data=xxxxxxxx")
	flag.Parse()

	payloads, err := parse.FromUri(*uri)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(payloads, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
