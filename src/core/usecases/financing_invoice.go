package usecases

import "encoding/json"

func FinancingInvoice(invoiceDTO map[string]interface{}) error {
	invoiceRepository.SaveInvoice(invoiceDTO)
	invoiceByte, _ := json.Marshal(invoiceDTO)
	messageBus.Produce(invoiceByte, "invoice_financed")
	return nil
}
