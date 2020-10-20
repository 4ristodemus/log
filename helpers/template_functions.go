package helpers

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

/*
 * Takes an entry path, parses the markdown, outputs an HTML template
 */
func RenderMarkdown(entryPath string) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	e, err := os.Executable()
	if err != nil {
		panic(err)
	}

	basePath := path.Dir(e)

	data, err := ioutil.ReadFile(
		path.Join(basePath, strings.Trim(entryPath, " \n")),
	)
	if err != nil {
		panic(err)
	}

	html := markdown.ToHTML(data, parser, nil)
	return template.HTML(string(html))
}

/*
 * Formats a go timestamp
 */
func FormatTimestamp(timestamp time.Time) string {
	return timestamp.Format("15:04 / Jan-2-2006")
}
