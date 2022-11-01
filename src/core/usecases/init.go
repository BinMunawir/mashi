package usecases

import services "github.com/BinMunawir/mashi/src/core/services/renderers"

var htmlRenderer services.HtmlRenderer

func Init(_htmlRenderer services.HtmlRenderer) {
	htmlRenderer = _htmlRenderer
}
