package links

import (
	"log/slog"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func LocalLinks(source []byte) (localDestinations []string) {
	md := goldmark.New()
	document := md.Parser().Parse(text.NewReader(source))
	ast.Walk(document, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if n.Kind() == ast.KindLink {
				link := n.(*ast.Link)
				var linkText string
				for child := link.FirstChild(); child != nil; child = child.NextSibling() {
					if text, ok := child.(*ast.Text); ok {
						linkText += string(text.Segment.Value(source))
					}
				}
				destination := string(link.Destination)
				slog.Info("link", "destination", destination, "linkText", linkText)
				if strings.HasPrefix(destination, "/") {
					// remove trailing slash
					destination = strings.TrimSuffix(destination, "/")
					// ignore all .jpg file links
					if !strings.Contains(destination, ".") {
						localDestinations = append(localDestinations, destination)
					}
				}
			}
		}
		return ast.WalkContinue, nil
	})
	return localDestinations
}
