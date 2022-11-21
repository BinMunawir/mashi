package usecases_test

import (
	"fmt"
	"testing"

	"github.com/BinMunawir/mashi/src/core/usecases"
)

func TestFinancingInvoice(t *testing.T) {
	initInvoiceRepositoryStub()

	var tests = []struct {
		input map[string]interface{}
	}{
		{map[string]interface{}{
			"sss": map[string]interface{}{
				"ddd": "dddd",
			},
		}},
	}
	for _, test := range tests {
		if got := usecases.FinancingInvoice(test.input); got != nil {
			t.Errorf("InvoiceFinancing(%v) = %v", test.input, got)
		}
	}
}

func initInvoiceRepositoryStub() {
	usecases.Init(invoiceRepositoryStub{})
}

type invoiceRepositoryStub struct{}

func (r invoiceRepositoryStub) SaveInvoice(invoiceDTO map[string]interface{}) {
	fmt.Println("SaveInvoice triggered with\n", invoiceDTO)
}
func (r invoiceRepositoryStub) RetrieveInvoice(id string) (invoice map[string]interface{}, err error) {
	return nil, nil
}
