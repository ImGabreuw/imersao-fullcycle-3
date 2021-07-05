package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAl       time.Time
}

func NewCreditCard() *CreditCard {
	creditCard := &CreditCard{}

	creditCard.ID = uuid.NewV4().String()
	creditCard.CreatedAl = time.Now()

	return creditCard
}
