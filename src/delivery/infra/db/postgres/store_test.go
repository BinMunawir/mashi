package postgres

import (
	"errors"
	"strings"
	"testing"
)

type PostgresStore struct{}

func NewPostgresStore(dsn string) (PostgresStore, error) {
	if !strings.Contains(dsn, "6") {
		return PostgresStore{}, errors.New("Unable to create connection")
	}
	return PostgresStore{}, nil
}
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
			in:   input{"postgres://mashi:123456789@db/mashi?sslmode=disable"},
			out:  output{PostgresStore{}, nil},
		},
		{
			name: "fail",
			in:   input{"postgres://mashi:12345789@db/mashi?sslmode=disable"},
			out:  output{PostgresStore{}, errors.New("Unable to create connection")},
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
