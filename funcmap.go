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
	"hkabb":     func() interface{} { return gofakeit.HackerAbbreviation() },
	"hkadj":     func() interface{} { return gofakeit.HackerAdjective() },
	"hknn":      func() interface{} { return gofakeit.HackerNoun() },
	"hkvb":      func() interface{} { return gofakeit.HackerVerb() },
	"hkph":      func() interface{} { return gofakeit.HackerPhrase() },
	"noun":      func() interface{} { return gofakeit.Noun() },
	"verb":      func() interface{} { return gofakeit.Verb() },
	"adverb":    func() interface{} { return gofakeit.Adverb() },
	"prep":      func() interface{} { return gofakeit.Preposition() },
	"adj":       func() interface{} { return gofakeit.Adjective() },
	"word":      func() interface{} { return gofakeit.Word() },
	"liword":    func() interface{} { return gofakeit.LoremIpsumWord() },
	"question":  func() interface{} { return gofakeit.Question() },
	"quote":     func() interface{} { return gofakeit.Quote() },
	"phrase":    func() interface{} { return gofakeit.Phrase() },

	// Custom generators
	"seq":    Sequence,
	"cardno": CardNumber,
	"uuid":   Uuid,
	"urn":    Urn,
	"card":   Card,
}

