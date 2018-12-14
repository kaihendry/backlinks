package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

func main() {
	sourceFilename := os.Args[1]
	b, err := ioutil.ReadFile(sourceFilename)
	if err != nil {
		panic(err)
	}
	links := re.FindAllStringSubmatch(string(b), -1)

	for _, l := range links {
		if len(l) != 3 {
			panic(fmt.Errorf("parsed link not in tuple: %v", l))
		}

		target := l[2]
		outputfile := strings.TrimSuffix(target, filepath.Ext(target)) + ".whatlinkshere"
		backLinkHTML := strings.TrimSuffix(sourceFilename, filepath.Ext(sourceFilename))
		log.Println(outputfile, backLinkHTML)

		// write the backlink to $target.whatlinkshere
		f, err := os.OpenFile(outputfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(fmt.Sprintf("[%s](%s)\n", backLinkHTML, backLinkHTML+".html")); err != nil {
			panic(err)
		}

	}
}
