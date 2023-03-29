package bill_test

import (
	"testing"

	"github.com/invopop/gobl/bill"
	"github.com/invopop/gobl/cal"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/l10n"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/regimes/co"
	"github.com/invopop/gobl/regimes/common"
	"github.com/invopop/gobl/tax"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvoiceCorrect(t *testing.T) {
	// Spanish Case (only corrective)
	i := testInvoiceESForCorrection(t)

	_, err := i.Correct(bill.Credit, bill.Debit)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot use both credit and debit options")

	i2, err := i.Correct(bill.Credit, bill.WithReason("test refund"))
	require.NoError(t, err)
	assert.Equal(t, bill.InvoiceTypeCorrective, i2.Type)
	assert.Equal(t, i2.Lines[0].Quantity.String(), "-10")
	pre := i2.Preceding[0]
	assert.Equal(t, pre.Series, i.Series)
	assert.Equal(t, pre.Code, i.Code)
	assert.Equal(t, pre.Reason, "test refund")

	err = i2.Calculate()
	require.NoError(t, err)

	_, err = i.Correct(bill.Debit, bill.WithReason("should fail"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "debit not supported")

	i2, err = i.Correct()
	require.NoError(t, err)
	assert.Equal(t, i2.Type, bill.InvoiceTypeCorrective)

	// France case (both corrective and credit note)
	i = testInvoiceFRForCorrection(t)
	i2, err = i.Correct()
	require.NoError(t, err)
	assert.Equal(t, i2.Type, bill.InvoiceTypeCorrective)

	i2, err = i.Correct(bill.Credit)
	require.NoError(t, err)
	assert.Equal(t, i2.Type, bill.InvoiceTypeCreditNote)

	// Colombia case (only credit note)

	i = testInvoiceCOForCorrection(t)
	_, err = i.Correct(bill.Credit)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing stamp")

	stamps := []*cbc.Stamp{
		{
			Provider: co.StampProviderDIANCUDE,
			Value:    "FOOO",
		},
		{
			Provider: co.StampProviderDIANQR, // not copied!
			Value:    "BARRRR",
		},
	}

	_, err = i.Correct(bill.WithStamps(stamps))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "correction not supported by regime")

	i2, err = i.Correct(bill.Credit, bill.WithStamps(stamps), bill.WithCorrectionMethod(co.CorrectionMethodKeyRevoked))
	require.NoError(t, err)
	assert.Equal(t, i2.Type, bill.InvoiceTypeCreditNote)
	pre = i2.Preceding[0]
	require.Len(t, pre.Stamps, 1)
	assert.Equal(t, pre.Stamps[0].Provider, co.StampProviderDIANCUDE)
	assert.Equal(t, pre.CorrectionMethod, co.CorrectionMethodKeyRevoked)
}

func testInvoiceESForCorrection(t *testing.T) *bill.Invoice {
	t.Helper()
	i := &bill.Invoice{
		Series: "TEST",
		Code:   "123",
		Tax: &bill.Tax{
			PricesInclude: common.TaxCategoryVAT,
		},
		Supplier: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.ES,
				Code:    "B98602642",
			},
		},
		Customer: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.ES,
				Code:    "54387763P",
			},
		},
		IssueDate: cal.MakeDate(2022, 6, 13),
		Lines: []*bill.Line{
			{
				Quantity: num.MakeAmount(10, 0),
				Item: &org.Item{
					Name:  "Test Item",
					Price: num.MakeAmount(10000, 2),
				},
				Taxes: tax.Set{
					{
						Category: "VAT",
						Rate:     "standard",
					},
				},
				Discounts: []*bill.LineDiscount{
					{
						Reason:  "Testing",
						Percent: num.NewPercentage(10, 2),
					},
				},
			},
		},
	}
	return i
}

func testInvoiceFRForCorrection(t *testing.T) *bill.Invoice {
	t.Helper()
	i := &bill.Invoice{
		Series: "TEST",
		Code:   "123",
		Supplier: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.FR,
				Code:    "732829320",
			},
		},
		Customer: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.FR,
				Code:    "391838042",
			},
		},
		IssueDate: cal.MakeDate(2022, 6, 13),
		Lines: []*bill.Line{
			{
				Quantity: num.MakeAmount(10, 0),
				Item: &org.Item{
					Name:  "Test Item",
					Price: num.MakeAmount(10000, 2),
				},
				Taxes: tax.Set{
					{
						Category: "VAT",
						Rate:     "standard",
					},
				},
				Discounts: []*bill.LineDiscount{
					{
						Reason:  "Testing",
						Percent: num.NewPercentage(10, 2),
					},
				},
			},
		},
	}
	return i
}

func testInvoiceCOForCorrection(t *testing.T) *bill.Invoice {
	t.Helper()
	i := &bill.Invoice{
		Series: "TEST",
		Code:   "123",
		Tax: &bill.Tax{
			PricesInclude: common.TaxCategoryVAT,
		},
		Supplier: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.CO,
				Code:    "9014586527",
			},
		},
		Customer: &org.Party{
			TaxID: &tax.Identity{
				Country: l10n.CO,
				Code:    "8001345363",
			},
		},
		IssueDate: cal.MakeDate(2022, 6, 13),
		Lines: []*bill.Line{
			{
				Quantity: num.MakeAmount(10, 0),
				Item: &org.Item{
					Name:  "Test Item",
					Price: num.MakeAmount(10000, 2),
				},
				Taxes: tax.Set{
					{
						Category: "VAT",
						Rate:     "standard",
					},
				},
				Discounts: []*bill.LineDiscount{
					{
						Reason:  "Testing",
						Percent: num.NewPercentage(10, 2),
					},
				},
			},
		},
	}
	return i
}
