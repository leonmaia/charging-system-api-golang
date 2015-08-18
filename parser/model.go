package parser

import (
	"errors"
	"time"
)

type Transaction struct {
	CustomerID string    `schema:"customerId"`
	StartTime  time.Time `schema:"startTime"`
	EndTime    time.Time `schema:"endTime"`
	Volume     float64   `schema:"volume"`
}

func (t Transaction) IsValid() error {
	if len(t.CustomerID) == 0 {
		return errors.New("Customer ID should not be empty")
	}

	return nil
}
