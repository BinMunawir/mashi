package repositories

type InvoiceRepository interface {
	SaveInvoice(map[string]interface{})
	RetrieveInvoice(id string) (invoice map[string]interface{}, err error)
}
