// Package facturae provides the FacturaE addon for Spanish invoices.
package facturae

import (
	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/tax"
)

const (
	// V3 for FacturaE versions 3.x
	V3 cbc.Key = "es-facturae-v3"
)

func init() {
	tax.RegisterAddonDef(newAddon())
}

func newAddon() *tax.AddonDef {
	return &tax.AddonDef{
		Key: V3,
		Name: i18n.String{
			i18n.EN: "Spain FacturaE",
		},
		Extensions:  extensions,
		Normalizer:  normalize,
		Scenarios:   scenarios,
		Validator:   validate,
		Corrections: invoiceCorrectionDefinitions,
	}
}

func normalize(doc any) {
	switch obj := doc.(type) {
	case *bill.Invoice:
		normalizeInvoice(obj)
	}
}

func validate(doc any) error {
	switch obj := doc.(type) {
	case *bill.Invoice:
		return validateInvoice(obj)
	}
	return nil
}
