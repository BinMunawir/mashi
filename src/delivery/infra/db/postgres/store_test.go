package postgres_test

import (
	"errors"
	"testing"

	"github.com/BinMunawir/mashi/src/delivery/infra/configs"
	"github.com/BinMunawir/mashi/src/delivery/infra/db/postgres"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgresStore(t *testing.T) {
	assert := assert.New(t)
	type input struct{ dsn string }
	type output struct {
		res postgres.PostgresStore
		err error
	}
	var tests = []struct {
		name        string
		in          input
		out         output
		extraAssert func(error)
	}{
		{
			"SuccessConnection",
			input{configs.DNS},
			output{postgres.PostgresStore{}, nil},
			func(err error) {
				assert.Nil(err)
			},
		},
		{
			"FailedConnection",
			input{"postgres://mashi:blablabla@0.0.0.0/mashi?sslmode=disable"},
			output{postgres.PostgresStore{}, errors.New("unable to create connection")},
			func(err error) {
				assert.NotNil(err)
				assert.ErrorContains(err, "unable to create connection")
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := postgres.NewPostgresStore(tc.in.dsn)

			tc.extraAssert(err)

		})
	}
}

func TestSaveInvoice(t *testing.T) {
	// assert := assert.New(t)
	type input struct{ data map[string]interface{} }
	type output struct{}
	var tests = []struct {
		name        string
		in          input
		out         output
		extraAssert func()
	}{
		{
			"empty",
			input{map[string]interface{}{
				"id":    "dkj-54123",
				"title": "dummy test invoice",
			}},
			output{},
			func() {
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			postgresStore, _ := postgres.NewPostgresStore(configs.DNS)
			invoiceStore := postgresStore
			invoiceStore.SaveInvoice(tc.in.data)

			tc.extraAssert()

			
		})
	}
}
