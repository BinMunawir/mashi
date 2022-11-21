package usecases

import (
	"github.com/BinMunawir/mashi/src/core/repositories"
	"github.com/BinMunawir/mashi/src/core/services"
)

var invoiceRepository repositories.InvoiceRepository
var messageBus services.MessageBus

func Init(_invoiceRepository repositories.InvoiceRepository, _messageBus services.MessageBus) {
	invoiceRepository = _invoiceRepository
	messageBus = _messageBus
}
