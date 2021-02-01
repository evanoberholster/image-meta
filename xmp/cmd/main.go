package main

import (
	"fmt"
	"os"

	"github.com/evanoberholster/imagemeta/xmp"
)

const (
	dir  = "../../test/img/"
	dir2 = "test/samples/"
	name = "CanonEOS7D.xmp"
)

func main() {
	xmp.DebugMode = true
	f, err := os.Open(dir2 + name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	xmp, err := xmp.Read(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xmp)
}