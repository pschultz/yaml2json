CLI interface for github.com/ghodss/yaml

### Installation

    go get github.com/pschultz/yaml2json

### Usage

    Usage: yaml2json <OPTIONS> <FILENAME>

    Convert the yaml contents of FILENAME to json and print it on stdout. If
    FILENAME is '-' or omitted, stdin is read instead.

    The input is read completely before re-encoding begins. Multiple yaml documents
    in the input are also not supported. yaml2json is thus of limited use in streaming
    pipelines.

    OPTIONS:
      -c    Print compact instead of pretty json.
