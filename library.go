package main

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"
)

func CardNumber() string {
	return gofakeit.CreditCard().Number
}

func Urn() string {
	guid, _ := uuid.NewUUID()
	return guid.URN()
}

func Uuid() string {
	guid, _ := uuid.NewUUID()
	return guid.String()
}
