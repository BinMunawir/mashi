package usecases

import (
	"reflect"
	"testing"
)

func TestSkeletonReport(t *testing.T) {
	var tests = []struct {
		out map[string]string
	}{
		{map[string]string{"content": "Skeleton report"}},
	}
	for _, tc := range tests {
		t.Run("SkeletonReport()", func(t *testing.T) {
			out, err := SkeletonReport()
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("SkeletonReport() = %v", out)
			}
		})
	}
	// for _, tc := range tests {
	// 	got, err := SkeletonReport()
	// 	if err != nil || !reflect.DeepEqual(got, tc.want) {
	// 		t.Errorf("SkeletonReport() = %v \terror: %v", got, err)
	// 	}
	// }
}
