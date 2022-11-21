package dtos

import "time"

type InvoiceDTO struct {
	Id        string
	Title     string
	Platform  string
	BeginDate time.Time
	DueDate   time.Time
	ROI       float32
}
