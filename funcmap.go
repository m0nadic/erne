package main

import gofakeit "github.com/brianvoe/gofakeit/v5"

type Generator func() interface{}

var genMap = map[string]Generator{
	"name":      func() interface{} { return gofakeit.Name() },
	"email":     func() interface{} { return gofakeit.Email() },
	"fruit":     func() interface{} { return gofakeit.Fruit() },
	"vegetable": func() interface{} { return gofakeit.Vegetable() },
	"breakfast": func() interface{} { return gofakeit.Breakfast() },
	"lunch":     func() interface{} { return gofakeit.Lunch() },
	"dinner":    func() interface{} { return gofakeit.Dinner() },
	"snack":     func() interface{} { return gofakeit.Snack() },
	"dessert":   func() interface{} { return gofakeit.Dessert() },
	"color":     func() interface{} { return gofakeit.Color() },
	"hexc":      func() interface{} { return gofakeit.HexColor() },
	"safec":     func() interface{} { return gofakeit.SafeColor() },
	"url":       func() interface{} { return gofakeit.URL() },
	"domnm":     func() interface{} { return gofakeit.DomainName() },
	"domsf":     func() interface{} { return gofakeit.DomainSuffix() },
	"ipv4":      func() interface{} { return gofakeit.IPv4Address() },
	"ipv6":      func() interface{} { return gofakeit.IPv6Address() },
	"ua":        func() interface{} { return gofakeit.UserAgent() },
	"uacr":      func() interface{} { return gofakeit.ChromeUserAgent() },
	"uaff":      func() interface{} { return gofakeit.FirefoxUserAgent() },
	"uaop":      func() interface{} { return gofakeit.OperaUserAgent() },
	"uasf":      func() interface{} { return gofakeit.SafariUserAgent() },
	"cardcvv":   func() interface{} { return gofakeit.CreditCardCvv() },
	"cardexp":   func() interface{} { return gofakeit.CreditCardExp() },
	"cardtype":  func() interface{} { return gofakeit.CreditCardType() },
	"seq":       Sequence,
	"cardno":    CardNumber,
	"uuid":      Uuid,
	"urn":       Urn,
	"card": Card,
}


func Card() interface{} {
	return map[string]string{
		"number": gofakeit.CreditCard().Number,
		"cvv": gofakeit.CreditCard().Cvv,
		"type": gofakeit.CreditCard().Type,
		"expiry": gofakeit.CreditCard().Exp,
	}
}
