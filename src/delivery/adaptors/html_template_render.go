package adaptors

import (
	"bytes"
	"html/template"
	"log"
)

func HTMLRender(data map[string]string) []byte {
	tmpl, err := template.ParseFiles("src/delivery/infra/templates/welcome_report.html")
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	return buf.Bytes()
}
