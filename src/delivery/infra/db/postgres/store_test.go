package postgres

import (
	"errors"
	"strings"
	"testing"

	"github.com/BinMunawir/mashi/src/delivery/infra/configs"
)

func TestNewPostgresStore(t *testing.T) {
	type input struct{ dsn string }
	type output struct {
		res PostgresStore
		err error
	}
	var tests = []struct {
		name string
		in   input
		out  output
	}{
		{
			name: "success",
			in:   input{configs.DNS},
			out:  output{PostgresStore{}, nil},
		},
		{
			name: "fail",
			in:   input{"postgres://mashi:blablabla@0.0.0.0/mashi?sslmode=disable"},
			out:  output{PostgresStore{}, errors.New("unable to create connection")},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPostgresStore(tc.in.dsn)
			if tc.out.err != nil && err == nil {
				t.Errorf("expect error: %v got nil error", tc.out.err)
			}
			if tc.out.err != nil && err != nil {
				if !strings.Contains(err.Error(), tc.out.err.Error()) {
					t.Errorf("expect error message to contains %v but got %v", tc.out.err.Error(), err.Error())
				}
			}
			if tc.out.err == nil && err != nil {
				t.Fatalf("Error: %v", err)
			}
		})
	}
}
