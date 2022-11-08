package usecases

import (
	"reflect"
	"testing"
)

func TestSkeletonReport(t *testing.T) {
	type input struct{}
	type output struct {
		res map[string]string
		err error
	}
	var tests = []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "default",
			in:   input{},
			out:  output{map[string]string{"content": "Skeleton report"}, nil},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SkeletonReport()
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.out.res) {
				t.Errorf("%v != %v", got, tc.out.res)
			}
		})
	}
}
