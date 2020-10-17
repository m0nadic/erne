package main

import "github.com/brianvoe/gofakeit/v5"

type Generator func() string

var genMap = map[string]Generator{
	"name":      gofakeit.Name,
	"email":     gofakeit.Email,
	"uuid":      Uuid,
	"urn":       Urn,
	"fruit":     gofakeit.Fruit,
	"vegetable": gofakeit.Vegetable,
	"breakfast": gofakeit.Breakfast,
	"lunch":     gofakeit.Lunch,
	"dinner":    gofakeit.Dinner,
	"snack":     gofakeit.Snack,
	"dessert":   gofakeit.Dessert,
	"color":     gofakeit.Color,
	"hexc":      gofakeit.HexColor,
	"safec":     gofakeit.SafeColor,
	"url":       gofakeit.URL,
	"domnm":     gofakeit.DomainName,
	"domsf":     gofakeit.DomainSuffix,
	"ipv4":      gofakeit.IPv4Address,
	"ipv6":      gofakeit.IPv6Address,
	"ua":        gofakeit.UserAgent,
	"uacr":      gofakeit.ChromeUserAgent,
	"uaff":      gofakeit.FirefoxUserAgent,
	"uaop":      gofakeit.OperaUserAgent,
	"uasf":      gofakeit.SafariUserAgent,
	"seq":       Sequence,
	"cardno":    CardNumber,
	"cardcvv":   gofakeit.CreditCardCvv,
	"cardexp":   gofakeit.CreditCardExp,
	"cardtype":  gofakeit.CreditCardType,
}
