package gr

import (
	"strings"

	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/head"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/pay"
	"github.com/invopop/gobl/tax"
	"github.com/invopop/validation"
)

// invoiceValidator adds validation checks to v.invoices which are relevant
// for the region.
type invoiceValidator struct {
	inv *bill.Invoice
}

func validateInvoice(inv *bill.Invoice) error {
	v := &invoiceValidator{inv: inv}
	return v.validate()
}

func (v *invoiceValidator) validate() error {
	return validation.ValidateStruct(v.inv,
		validation.Field(&v.inv.Series, validation.Required),
		validation.Field(&v.inv.Tax,
			validation.Required,
			validation.By(v.validateTax),
			validation.Skip,
		),
		validation.Field(&v.inv.Supplier,
			validation.By(v.validateBusinessParty),
			validation.Skip,
		),
		validation.Field(&v.inv.Customer,
			validation.When(!IsRetail(v.inv),
				validation.Required,
				validation.By(v.validateBusinessParty),
				validation.By(v.validateBusinessCustomer),
			),
			validation.Skip,
		),
		validation.Field(&v.inv.Lines,
			validation.Each(
				validation.By(v.validateLine),
				validation.Skip,
			),
			validation.Skip,
		),
		validation.Field(&v.inv.Payment,
			validation.Required,
			validation.By(v.validatePayment),
			validation.Skip,
		),
		validation.Field(&v.inv.Preceding,
			validation.When(
				v.inv.Type.In(bill.InvoiceTypeCreditNote),
				validation.Required,
			),
			validation.Each(validation.By(v.validatePreceding)),
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validateTax(value any) error {
	t, ok := value.(*bill.Tax)
	if !ok || t == nil {
		return nil
	}
	return validation.ValidateStruct(t,
		validation.Field(&t.Ext,
			tax.ExtensionsRequires(ExtKeyMyDATAInvoiceType),
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validateBusinessParty(value any) error {
	p, ok := value.(*org.Party)
	if !ok || p == nil {
		return nil
	}
	return validation.ValidateStruct(p,
		validation.Field(&p.TaxID,
			validation.Required,
			tax.RequireIdentityCode,
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validateBusinessCustomer(value any) error {
	p, ok := value.(*org.Party)
	if !ok || p == nil {
		return nil
	}
	return validation.ValidateStruct(p,
		validation.Field(&p.Addresses,
			validation.Required,
			validation.Length(1, 0),
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validateLine(value any) error {
	l, ok := value.(*bill.Line)
	if !ok || l == nil {
		return nil
	}
	return validation.ValidateStruct(l,
		validation.Field(&l.Total,
			num.Positive,
			num.NotZero,
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validatePayment(value any) error {
	p, ok := value.(*bill.Payment)
	if !ok || p == nil {
		return nil
	}
	return validation.ValidateStruct(p,
		validation.Field(&p.Instructions,
			validation.Required,
			validation.By(v.validatePaymentInstructions),
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validatePaymentInstructions(value any) error {
	i, ok := value.(*pay.Instructions)
	if !ok || i == nil {
		return nil
	}

	return validation.ValidateStruct(i,
		validation.Field(&i.Key,
			validation.Required,
			isValidPaymentMeanKey,
			validation.Skip,
		),
	)
}

func (v *invoiceValidator) validatePreceding(value any) error {
	p, ok := value.(*bill.Preceding)
	if !ok || p == nil {
		return nil
	}

	return validation.ValidateStruct(p,
		validation.Field(&p.Stamps,
			head.StampsHas(StampIAPRMark),
			validation.Skip,
		),
	)
}

func validateTaxCombo(tc *tax.Combo) error {
	return validation.ValidateStruct(tc,
		validation.Field(&tc.Ext,
			tax.ExtensionsRequires(ExtKeyMyDATAVATCat),
			validation.When(
				tc.Percent == nil,
				tax.ExtensionsRequires(ExtKeyMyDATAExemption),
			),
			validation.When(
				tc.Ext.Has(ExtKeyMyDATAIncomeCat) || tc.Ext.Has(ExtKeyMyDATAIncomeType),
				tax.ExtensionsRequires(ExtKeyMyDATAIncomeCat, ExtKeyMyDATAIncomeType),
			),
			validation.Skip,
		),
	)
}

// IsRetail returns true if the invoice type corresponds to a retail invoice.
func IsRetail(inv *bill.Invoice) bool {
	if inv.Tax == nil || inv.Tax.Ext == nil {
		return false
	}

	it := inv.Tax.Ext[ExtKeyMyDATAInvoiceType]

	return strings.HasPrefix(string(it), InvoiceTypeRetailPrefix)
}
