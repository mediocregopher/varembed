package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	pkg     = flag.String("pkg", "", "Name of the package the file will be embedded into")
	in      = flag.String("in", "", "Name of the file to read in")
	out     = flag.String("out", "", "Name of the go source file to create")
	varname = flag.String("varname", "", "Name of the variable to assign the file's contents to")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*in)
	if err != nil {
		panic(err)
	}

	b64 := base64.StdEncoding.EncodeToString(b)

	outf, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	_, err = fmt.Fprintf(outf, `package %s

import "encoding/base64"

var %s, _ = base64.StdEncoding.DecodeString(%q)
`,
		*pkg, *varname, b64)
	if err != nil {
		panic(err)
	}
}
