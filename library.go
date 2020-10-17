package main

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"
)

func CardNumber() interface{} {
	return gofakeit.CreditCard().Number
}

func Urn() interface{} {
	guid, _ := uuid.NewUUID()
	return guid.URN()
}

func Uuid() interface{} {
	guid, _ := uuid.NewUUID()
	return guid.String()
}

func Card() interface{} {
	return map[string]string{
		"number": gofakeit.CreditCard().Number,
		"cvv": gofakeit.CreditCard().Cvv,
		"type": gofakeit.CreditCard().Type,
		"expiry": gofakeit.CreditCard().Exp,
	}
}