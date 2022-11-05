package usecases

func InvoiceFinancing(invoiceDTO map[string]interface{}) error {
	invoiceRepository.SaveInvoice(invoiceDTO)
	return nil
}
