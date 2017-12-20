package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

func main() {
	prog := filepath.Base(os.Args[0])

	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s <OPTIONS> <FILENAME>

Convert the YAML contents of FILENAME to JSON and print it on stdout. If
FILENAME is '-' or omitted, stdin is read instead.

The input is read completely before re-encoding begins. Multiple YAML documents
in the input are also not supported. %s is thus of limited use in streaming
pipelines.

OPTIONS:
`, prog, prog)
		flag.PrintDefaults()
	}

	compact := false
	flag.BoolVar(&compact, "c", false, "Print compact instead of pretty JSON.")

	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		log.Fatalf("Expected at most one argument; the name of a YAML file")
	}

	b, err := readInput(args)
	check(err)

	b, err = yaml.YAMLToJSON(b)
	check(err)

	if compact {
		os.Stdout.Write(append(b, '\n'))
		return
	}

	b, err = makePretty(b)
	check(err)

	os.Stdout.Write(append(b, '\n'))
}

func readInput(args []string) ([]byte, error) {
	if len(args) > 0 && args[0] != "-" {
		return ioutil.ReadFile(args[0])
	} else {
		return ioutil.ReadAll(os.Stdin)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makePretty(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.Indent(buf, b, "", "  ")
	return buf.Bytes(), err
}
