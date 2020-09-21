package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Le0tk0k/imgconv"
)

var from = flag.String("f", ".jpeg", "Extension before conversion")
var to = flag.String("t", ".png", "Extension after conversion")
var rm = flag.Bool("r", false, "Remove file before conversion")

func main() {
	flag.Parse()
	src := flag.Arg(0)
	dst := strings.Replace(flag.Arg(0), *from, *to, 1)

	err := imgconv.Convert(src, dst, *rm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
