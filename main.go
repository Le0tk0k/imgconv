package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Le0tk0k/imgpro/imgconv"
)

var from = flag.String("from", ".jpeg", "Extension before conversion")
var to = flag.String("to", ".png", "Extension after conversion")

func main() {
	flag.Parse()
	err := imgconv.Convert(flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
