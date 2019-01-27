package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/dihedron/go-log"
	"github.com/dihedron/jsonx/jsonx"
)

func init() {
	switch strings.ToLower(os.Getenv("JSONX_DEBUG")) {
	case "debug", "dbg", "d":
		log.SetLevel(log.DBG)
	case "informational", "information", "info", "inf", "i":
		log.SetLevel(log.INF)
	case "warning", "warn", "wrn", "w":
		log.SetLevel(log.WRN)
	case "error", "err", "e":
		log.SetLevel(log.ERR)
	default:
		log.SetLevel(log.NUL)
	}
	log.SetStream(os.Stderr, true)
	log.SetTimeFormat("15:04:05.000")
	log.SetPrintCallerInfo(true)
	log.SetPrintSourceInfo(log.SourceInfoShort)
}

func main() {

	help := flag.Bool("help", false, "prints help information and quits")
	flag.Parse()
	if *help {
		fmt.Fprintf(os.Stderr, `
 JSON does not admit comments; anyway being able to include free text that is 
 ignored by JSON parsers would be extremely useful, especially in configuration 
 files or in any other source that is used to store human-readable information.
 JSONX is meant as a way to remove C++-stype (//...) and shell-style (#...) 
 // comments from JSON-like files; it can be used to read from a file or from 
 // STDIN and to write to a file or STDOUT.

 usage:
	 jsonx [<input> [<output>]]
   that is, 
	 if no arguments are provided it will read from STDIN and write to STDOUT
	 if a single argument is provided, it must be a valid file; output is STDOUT
	 if two arguments are provided, first is the input and the second the output

 examples:
     $> cat myfile.jsonx | jsonx > myfile.json
`)
		os.Exit(1)
	}

	input, err := getInput(os.Args)
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer input.Close()

	output, err := getOutput(os.Args)
	if err != nil {
		log.Fatalf("unable to open output file: %v", err)
	}
	defer output.Close()

	jsonx.Parse(input, output)
}

// getInput returns the input Reader to use; if a filename argument is provided,
// open the file to read from it, otherwise return STDIN; the Reader must be
// closed by the method's caller.
func getInput(args []string) (*os.File, error) {
	if len(args) > 1 {
		log.Debugf("reading text from input file: %q", args[1])
		return os.Open(args[1])
	}
	return os.Stdin, nil
}

// getOutput returns the output Writer to use; if a filename argument is provided,
// open the file to write to it, otherwise return STDOUT; the Writer must be
// closed by the method's caller.
func getOutput(args []string) (*os.File, error) {
	if len(args) > 2 {
		log.Debugf("writing text to output file: %q", args[2])
		return os.Create(args[2])
	}
	return os.Stdout, nil
}
