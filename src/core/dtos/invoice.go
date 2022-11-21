package dtos

import "time"

type Invoice struct {
	Id        string
	Title     string
	Platform  string
	BeginDate time.Time
	DueDate   time.Time
	ROI       int
}
