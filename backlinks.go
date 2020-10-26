package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	sourceFilename := os.Args[1]
	if filepath.Ext(sourceFilename) != ".mdwn" {
		fmt.Fprintf(os.Stderr, "source %s is not markdown\n", sourceFilename)
		os.Exit(1)
	}
	log.Println("Analysing", sourceFilename, "for new links")

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
		outputfile := strings.TrimSuffix(target, filepath.Ext(target)) + ".bl"
		backLinkHTML := strings.TrimSuffix(sourceFilename, filepath.Ext(sourceFilename))
		log.Println(outputfile, backLinkHTML)

		lines, err := readLines(outputfile)
		if err != nil {
			log.Printf("Initialising: %s", outputfile)
		}

		backlink := fmt.Sprintf("[%s](%s)", backLinkHTML, backLinkHTML+".html")

		var seenbefore bool
		for _, line := range lines {
			if backlink == line {
				seenbefore = true
			}
		}
		if !seenbefore {
			lines = append(lines, backlink)
			log.Printf("Wrote: %s into %s", backlink, outputfile)
			if err := writeLines(lines, outputfile); err != nil {
				log.Fatalf("writeLines: %s", err)
			}
		}
	}
}
