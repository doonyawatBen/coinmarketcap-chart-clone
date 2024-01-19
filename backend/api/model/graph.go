package model

import "time"

type RequestGraph struct {
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate   time.Time `json:"endDate" validate:"required"`
	Width     string    `json:"width" validate:"required"`
	Height    string    `json:"height" validate:"required"`
}
