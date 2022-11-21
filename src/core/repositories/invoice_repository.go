package repositories

import "github.com/BinMunawir/mashi/src/core/dtos"

type InvoiceRepository interface {
	SaveInvoice(dtos.InvoiceDTO)
	RetrieveInvoice(id string) (invoice map[string]interface{}, err error)
}
