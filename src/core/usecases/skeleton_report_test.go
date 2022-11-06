package usecases

import (
	"reflect"
	"testing"
)

func TestSkeletonReport(t *testing.T) {
	type output struct {
		res map[string]string
		err error
	}
	type input struct{}
	var tests = []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "default",
			in:   input{},
			out:  output{map[string]string{"content": "Skeletoln report"}, nil},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := output{}
			out.res, out.err = SkeletonReport()
			if out.err != nil {
				t.Fatalf("Error: %v", out.err)
			}
			if !reflect.DeepEqual(out.res, tc.out.res) {
				t.Errorf("%v != %v", out.res, tc.out.res)
			}
		})
	}
}
