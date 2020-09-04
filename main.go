package main

import (
	"flag"
)

var from = flag.String("from", ".jpeg", "Extension before conversion")
var to = flag.String("to", ".png", "Extension after conversion")

func main() {
	flag.Parse()
}
