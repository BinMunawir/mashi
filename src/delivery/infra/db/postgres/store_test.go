package postgres_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/BinMunawir/mashi/src/core/dtos"
	"github.com/BinMunawir/mashi/src/core/repositories"
	"github.com/BinMunawir/mashi/src/delivery/infra/configs"
	"github.com/BinMunawir/mashi/src/delivery/infra/db/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
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
	assert := assert.New(t)
	type input struct{ invoice dtos.InvoiceDTO }
	type output struct{}
	var tests = []struct {
		name        string
		in          input
		out         output
		extraAssert func()
	}{
		{
			"empty",
			input{dtos.InvoiceDTO{
				Id:        "id",
				Title:     "title",
				Platform:  "platform",
				BeginDate: time.Now().UTC().Add(time.Duration(time.Now().Day())),
				DueDate:   time.Now(),
				ROI:       11.90,
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
			invoiceStore.SaveInvoice(tc.in.invoice)

			tc.extraAssert()

			pool, _ := pgxpool.New(context.Background(), configs.DNS)
			var invoice dtos.InvoiceDTO
			err := pool.QueryRow(context.Background(), "select id, title, platform, begin_date, due_date, roi from invoices;").
				Scan(&invoice.Id, &invoice.Title, &invoice.Platform, &invoice.BeginDate, &invoice.DueDate, &invoice.ROI)
			assert.Nil(err)

			assert.Equal(tc.in.invoice, invoice)

		})
	}
}

func TestRetrieveInvoice(t *testing.T) {
	assert := assert.New(t)
	type input struct{ id string }
	type output struct {
		invoice dtos.InvoiceDTO
		err     error
	}
	var tests = []struct {
		name        string
		in          input
		out         output
		extraAssert func()
	}{
		{
			"empty",
			input{"dkj-54123"},
			output{
				dtos.InvoiceDTO{
					Id:        "id",
					Title:     "title",
					Platform:  "platform",
					BeginDate: time.Now().UTC().Add(time.Duration(time.Now().Day())),
					DueDate:   time.Now(),
					ROI:       11.90,
				},
				nil,
			},
			func() {
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			postgresStore, _ := postgres.NewPostgresStore(configs.DNS)
			var invoiceStore repositories.InvoiceRepository = postgresStore
			invoiceStore.SaveInvoice(tc.out.invoice)
			invoice, err := invoiceStore.RetrieveInvoice(tc.in.id)

			tc.extraAssert()

			assert.Nil(err)
			assert.Equal(invoice, tc.out.invoice)

		})
	}
}
