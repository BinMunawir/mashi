package services

type HtmlRenderer interface {
	HTMLRender(tmplPath string, data map[string]string) []byte
}
