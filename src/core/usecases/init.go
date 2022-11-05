package usecases

import "github.com/BinMunawir/mashi/src/core/repositories"

var invoiceRepository repositories.InvoiceRepository

func Init(_invoiceRepository repositories.InvoiceRepository) {
	invoiceRepository = _invoiceRepository
}
