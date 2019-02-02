package jsonx

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	log "github.com/dihedron/go-log"
)

// Parse reads a JSONX file line by line, removing C-style (/*...*/), C++-style
// (//...) and shell-style (#...) comments.
func Parse(input, output *os.File) {
	scanner := bufio.NewScanner(input)
	var cppStyleComment = regexp.MustCompile(`(.*)(//.*)$`)
	var shellStyleComment = regexp.MustCompile(`(.*)(#.*)$`)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debugf("read line: %q\n", line)
		matches := cppStyleComment.FindStringSubmatch(line)
		if matches != nil {
			log.Debugf("line has a C++-style comment:\n - text  : %q\n - comment:%q", matches[1], matches[2])
			if len(strings.TrimSpace(matches[1])) > 0 {
				fmt.Fprintf(output, "%s\n", matches[1])
			}
			continue
		}
		matches = shellStyleComment.FindStringSubmatch(line)
		if matches != nil {
			log.Debugf("line has a shell-style comment:\n - text  : %q\n - comment:%q", matches[1], matches[2])
			if len(strings.TrimSpace(matches[1])) > 0 {
				fmt.Fprintf(output, "%s\n", matches[1])
			}
			continue
		}
		fmt.Fprintf(output, "%s\n", line)
	}
}
