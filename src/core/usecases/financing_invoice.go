package usecases

func FinancingInvoice(invoiceDTO map[string]interface{}) error {
	invoiceRepository.SaveInvoice(invoiceDTO)
	return nil
}
