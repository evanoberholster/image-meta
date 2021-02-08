package main

import (
	"fmt"
	"os"

	"github.com/evanoberholster/imagemeta/xmp"
)

const (
	dir   = "../../test/img/"
	dir2  = "test/samples/"
	name  = "CanonEOS7DII.xmp"
	name1 = "image1.jpeg.xmp"
)

func main() {
	xmp.DebugMode = true
	f, err := os.Open(dir2 + name) //"retouch.xmp") //name)
	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Println(f.Close())
	}()

	xmp, err := xmp.Read(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xmp)
}
