package usecases

import "encoding/json"

const tmplPath string = "src/delivery/infra/templates/skeleton_report.html"

func skeletonReportContent() map[string]string {
	return map[string]string{
		"content": "Skeleton report",
	}
}

func SkeletonReportHtml() []byte {
	data := skeletonReportContent()
	return htmlRenderer.RenderHtml(tmplPath, data)
}

func SkeletonReportJson() []byte {
	data := skeletonReportContent()
	json, _ := json.Marshal(data)
	return json
}
