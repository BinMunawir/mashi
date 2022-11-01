package adaptors

import (
	"bytes"
	"html/template"
	"log"

	services "github.com/BinMunawir/mashi/src/core/services/renderers"
)

type golangHtmlRenderer struct{}

func NewGolangHtmlRenderer() services.HtmlRenderer {
	return golangHtmlRenderer{}
}

func (r golangHtmlRenderer) RenderHtml(tmplPath string, data map[string]string) []byte {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	return buf.Bytes()
}
