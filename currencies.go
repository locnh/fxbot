package main

// validCurrencies contains a set of supported currency codes based on ISO 4217.
// This list is used to validate user input before making API calls.
var revolutHoldingCurrencies = map[string]bool{
	"AED": true, // United Arab Emirates Dirham
	"AUD": true, // Australian Dollar
	"BGN": true, // Bulgarian Lev
	"CAD": true, // Canadian Dollar
	"CHF": true, // Swiss Franc
	"CLP": true, // Chilean Peso
	"CNY": true, // Chinese Yuan Renminbi
	"COP": true, // Colombian Peso
	"CZK": true, // Czech Koruna
	"DKK": true, // Danish Krone
	"EGP": true, // Egyptian Pound
	"EUR": true, // Euro
	"GBP": true, // British Pound Sterling
	"HKD": true, // Hong Kong Dollar
	"HUF": true, // Hungarian Forint
	"IDR": true, // Indonesian Rupiah
	"ILS": true, // Israeli New Shekel
	"INR": true, // Indian Rupee
	"ISK": true, // Icelandic Krona
	"JPY": true, // Japanese Yen
	"KRW": true, // South Korean Won
	"KZT": true, // Kazakhstani Tenge
	"MAD": true, // Moroccan Dirham
	"MXN": true, // Mexican Peso
	"NOK": true, // Norwegian Krone
	"NZD": true, // New Zealand Dollar
	"PHP": true, // Philippine Peso
	"PLN": true, // Polish Zloty
	"QAR": true, // Qatari Riyal
	"RON": true, // Romanian Leu
	"RSD": true, // Serbian Dinar
	"SAR": true, // Saudi Riyal
	"SEK": true, // Swedish Krona
	"SGD": true, // Singapore Dollar
	"THB": true, // Thai Baht
	"TRY": true, // Turkish Lira
	"USD": true, // United States Dollar
	"VND": true, // Vietnamese Dong
	"ZAR": true, // South African Rand
}
