package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Le0tk0k/imgpro/imgconv"
)

var from = flag.String("from", ".jpeg", "Extension before conversion")
var to = flag.String("to", ".png", "Extension after conversion")
var rmsrc = flag.Bool("rmsrc", false, "Remove file before conversion")

func main() {
	flag.Parse()
	src := flag.Arg(0)
	dst := strings.Replace(flag.Arg(0), *from, *to, 1)

	err := imgconv.Convert(src, dst, *rmsrc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
