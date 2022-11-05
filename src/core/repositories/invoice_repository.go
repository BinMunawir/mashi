package repositories

type InvoiceRepository interface {
	SaveInvoice(map[string]interface{})
}
