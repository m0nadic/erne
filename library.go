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
