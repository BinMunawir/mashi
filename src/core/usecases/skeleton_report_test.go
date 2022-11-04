package usecases

import (
	"reflect"
	"testing"
)

func TestSkeletonReport(t *testing.T) {
	var tests = []struct {
		want map[string]string
	}{
		{map[string]string{"content": "Skeleton report"}},
	}
	for _, test := range tests {
		if got := SkeletonReport(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("SkeletonReport() = %v", got)
		}
	}
}
