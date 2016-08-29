package main

import (
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

Convert the yaml contents of FILENAME to json and print it on stdout. If
FILENAME is '-' or omitted, stdin is read instead.

The input is read completely before re-encoding begins. Multiple yaml documents
in the input are also not supported. %s is thus of limited use in streaming
pipelines.

OPTIONS:
`, prog, prog)
		flag.PrintDefaults()
	}

	compact := flag.Bool("c", false, "Print compact instead of pretty json.")

	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		log.Fatalf("expected at most one argument; the name of a yaml file")
	}

	b, err := readInput(args)
	check(err)

	b, err = yaml.YAMLToJSON(b)
	check(err)

	if !*compact {
		b, err = makePretty(b)
		check(err)
	}

	os.Stdout.Write(b)
	os.Stdout.Write([]byte{'\n'})
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
	var doc *json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return b, err
	}

	return json.MarshalIndent(doc, "", "  ")
}
