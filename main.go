package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(filepath.Base(os.Args[0]) + ": ")

	b, err := ioutil.ReadAll(os.Stdin)
	check(err)

	b, err = yaml.YAMLToJSON(b)
	check(err)

	var doc *json.RawMessage
	err = json.Unmarshal(b, &doc)
	check(err)

	b, err = json.MarshalIndent(doc, "", "  ")
	check(err)

	fmt.Fprintf(os.Stdout, "%s\n", b)
}
