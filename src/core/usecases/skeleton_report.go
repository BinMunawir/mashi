package usecases

const tmplPath string = "src/delivery/infra/templates/skeleton_report.html"

func SkeletonReportData() map[string]string {
	return map[string]string{
		"content": "Skeleton report",
	}
}

func SkeletonReportHtml() []byte {
	data := SkeletonReportData()
	return htmlRenderer.RenderHtml(tmplPath, data)
}
