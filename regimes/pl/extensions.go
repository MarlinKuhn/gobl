package pl

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
)

// Regime extension codes for local electronic formats.
const (
	ExtKeyKSeFVATZero       = "pl-ksef-vat-zero"
	ExtKeyKSeFVATSpecial    = "pl-ksef-vat-special"
	ExtKeyKSeFEffectiveDate = "pl-ksef-effective-date"
)

var extensionKeys = []*cbc.Definition{
	{
		Key: ExtKeyKSeFVATSpecial,
		Name: i18n.String{
			i18n.EN: "Special VAT Extensions for KSeF",
			i18n.PL: "Rozszerzenia specjalne dla KSeF",
		},
		Values: []*cbc.Definition{
			{
				Code: "taxi",
				Name: i18n.String{
					i18n.EN: "Taxi Rate",
					i18n.PL: "Ryczałt dla taksówek",
				},
				Desc: i18n.String{
					i18n.EN: "Special flat rate for taxi drivers.",
					i18n.PL: "Specjalna stawka ryczałtu dla taksówkarzy.",
				},
			},
		},
	},
	{
		Key: ExtKeyKSeFVATZero,
		Name: i18n.String{
			i18n.EN: "Zero VAT Extensions for KSeF",
		},
		Values: []*cbc.Definition{
			{
				Code: "wdt",
				Name: i18n.String{
					i18n.EN: "WDT",
					i18n.PL: "WDT",
				},
				// TODO: description required
			},
			{
				Code: "domestic",
				Name: i18n.String{
					i18n.EN: "Domestic",
					i18n.PL: "Krajowy",
				},
				// TODO: description required
			},
			{
				Code: "export",
				Name: i18n.String{
					i18n.EN: "Export",
					i18n.PL: "Eksport",
				},
				// TODO: description required
			},
		},
	},
	{
		Key: ExtKeyKSeFEffectiveDate,
		Name: i18n.String{
			i18n.EN: "Effective date code.",
			i18n.PL: "Kod daty wejścia w życie.",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Original",
					i18n.PL: "Pierwotna",
				},
				Desc: i18n.String{
					i18n.EN: "Effective according to date of the original invoice.",
					i18n.PL: "Faktura skutkująca w dacie ujęcia faktury pierwotnej.",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Correction",
					i18n.PL: "Korygująca",
				},
				Desc: i18n.String{
					i18n.EN: "Effective according to date of correction.",
					i18n.PL: "Faktura skutkująca w dacie ujęcia faktury korygującej.",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Other",
					i18n.PL: "Inna",
				},
				Desc: i18n.String{
					i18n.EN: "Correction has legal consequences in another date or the dates are different for different position on the invoice",
					i18n.PL: "Faktura skutkująca w innej dacie. W tym gdy dla różnych pozycji faktury korygującej data jest różna.",
				},
			},
		},
	},
}
