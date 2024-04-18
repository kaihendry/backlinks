package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaihendry/backlinks/links"
	"github.com/yuin/goldmark"
)

type Page struct {
	Content   template.HTML // rendered markdown
	Path      Path          // path to this page
	Links     []Path        // path pages that this page links to
	Backlinks []Path        // path pages that link to this page
}

type Path string // relative to the root, path to the page, e.g. pageA.mdwn = /pageA

func main() {
	// find *.mdwn files in the current directory, render them to out/ directory
	files, err := filepath.Glob("*.mdwn")
	if err != nil {
		log.Fatal(err)
	}
	pages, err := fetch(files)
	if err != nil {
		log.Fatal(err)
	}

	// for each page, work out which other pages link to it, and then set that as the backlinks
	// create a map of pages to backlinks
	backlinks := make(map[Path][]Path)
	for _, page := range pages {
		for _, link := range page.Links {
			slog.Info("found link", "page", page.Path, "link", link)
			backlinks[link] = append(backlinks[link], page.Path)
		}
	}
	for i, page := range backlinks {
		slog.Info("backlinks", "i", i, "linksfrom", page)
	}

	// set the backlinks on each page
	for i := 0; i < len(pages); i++ {
		if backlinks[pages[i].Path] != nil {
			slog.Info("found backlinks", "page", pages[i].Path, "backlinks", backlinks[pages[i].Path])
			pages[i].Backlinks = append(pages[i].Backlinks, backlinks[pages[i].Path]...)
		}
	}

	slog.Info("pages", "pages", pages)

	err = render(pages)
	if err != nil {
		log.Fatal(err)
	}
}

func render(pages []Page) error {
	destDir := "_site"
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.Mkdir(destDir, 0755)
	}
	var tmplFile = "index.gohtml"
	if _, err := os.Stat(tmplFile); os.IsNotExist(err) {
		return err
	}
	for _, page := range pages {
		t := template.Must(template.New(tmplFile).ParseFiles(tmplFile))
		destDir := filepath.Join(destDir, string(page.Path))
		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			return err
		}
		outFile := fmt.Sprintf("%s/%s", destDir, "index.html")
		slog.Info("writing", "destDir", destDir, "outFile", outFile)
		index, err := os.Create(outFile)
		if err != nil {
			slog.Error("failed to create index file", "err", err)
			return err
		}
		defer index.Close()
		t.Execute(index, page)
	}
	return nil
}

func fetch(files []string) (pages []Page, err error) {
	pages = make([]Page, 0, len(files))
	for _, file := range files {
		page, err := parse(file)
		if err != nil {
			return pages, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func path(file string) string {
	// pageA.mdwn -> /pageA
	return "/" + strings.TrimSuffix(file, filepath.Ext(file))
}

func parse(file string) (Page, error) {
	page := Page{Path: Path(path(file))}
	content, err := os.ReadFile(file)
	if err != nil {
		return page, err
	}
	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		panic(err)
	}
	page.Links = make([]Path, len(links.LocalLinks(content)))
	for i, link := range links.LocalLinks(content) {
		page.Links[i] = Path(link)
	}
	page.Content = template.HTML(buf.String())
	return page, nil
}
