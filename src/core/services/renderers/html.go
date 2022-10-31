package services

type HtmlRenderer interface {
	RenderHtml(tmplPath string, data map[string]string) []byte
}
