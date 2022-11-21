package usecases_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/BinMunawir/mashi/src/core/dtos"
	"github.com/BinMunawir/mashi/src/core/usecases"
	"github.com/stretchr/testify/assert"
)

// DOTO: not thread safe!
var (
	invoiceFinancedEventProduced = false
)

func TestFinancingInvoice(t *testing.T) {
	assert := assert.New(t)

	initStubs()

	type input struct{ invoice dtos.InvoiceDTO }
	type output struct{}
	var tests = []struct {
		name        string
		in          input
		out         output
		extraAssert func(error)
	}{
		{
			"FinancingInvoice",
			input{dtos.InvoiceDTO{
				Id:        "id",
				Title:     "title",
				Platform:  "platform",
				BeginDate: time.Now().UTC().Add(time.Duration(time.Now().Day())),
				DueDate:   time.Now(),
				ROI:       11.90,
			}},
			output{},
			func(err error) {},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := usecases.FinancingInvoice(tc.in.invoice)
			tc.extraAssert(err)
			assert.True(invoiceFinancedEventProduced)
		})
	}
}

func initStubs() {
	usecases.Init(invoiceRepositoryStub{}, messagebusStub{})
}

type invoiceRepositoryStub struct{}

func (r invoiceRepositoryStub) SaveInvoice(invoice dtos.InvoiceDTO) {
	fmt.Println("SaveInvoice triggered with\n", invoice)
}
func (r invoiceRepositoryStub) RetrieveInvoice(id string) (invoice map[string]interface{}, err error) {
	return nil, nil
}

type messagebusStub struct{}

func (b messagebusStub) Produce(m []byte, topic string) error {
	invoiceFinancedEventProduced = true
	return nil
}
func (b messagebusStub) Consume(topic string) ([]byte, error) {
	return nil, nil
}
