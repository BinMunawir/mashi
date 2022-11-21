package usecases

import (
	"encoding/json"

	"github.com/BinMunawir/mashi/src/core/dtos"
)

func FinancingInvoice(invoice dtos.InvoiceDTO) error {
	invoiceRepository.SaveInvoice(invoice)
	invoiceByte, _ := json.Marshal(invoice)
	messageBus.Produce(invoiceByte, "invoice_financed")
	return nil
}
