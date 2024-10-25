package en16931_test

import (
	"testing"

	"github.com/invopop/gobl/addons/eu/en16931"
	"github.com/invopop/gobl/catalogues/untdid"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/regimes/es"
	"github.com/invopop/gobl/tax"
	"github.com/stretchr/testify/assert"
)

func TestTaxComboNormalization(t *testing.T) {
	ad := tax.AddonForKey(en16931.V2017)
	t.Run("standard VAT rate", func(t *testing.T) {
		p := num.MakePercentage(19, 2)
		c := &tax.Combo{
			Category: tax.CategoryVAT,
			Rate:     tax.RateStandard,
			Percent:  &p,
		}
		ad.Normalizer(c)
		assert.Equal(t, "S", c.Ext[untdid.ExtKeyTaxCategory].String())
		assert.Equal(t, "19%", c.Percent.String())
	})

	t.Run("unkown rate", func(t *testing.T) {
		c := &tax.Combo{
			Category: tax.CategoryVAT,
			Rate:     cbc.Key("unknown"),
			Percent:  num.NewPercentage(19, 2),
		}
		ad.Normalizer(c)
		assert.Empty(t, c.Ext)
	})
	t.Run("IGIC", func(t *testing.T) {
		c := &tax.Combo{
			Category: es.TaxCategoryIGIC,
			Percent:  num.NewPercentage(7, 2),
		}
		ad.Normalizer(c)
		assert.Equal(t, "L", c.Ext[untdid.ExtKeyTaxCategory].String())
		assert.Equal(t, "7%", c.Percent.String())
	})

	t.Run("IPSI", func(t *testing.T) {
		c := &tax.Combo{
			Category: es.TaxCategoryIPSI,
			Percent:  num.NewPercentage(7, 2),
		}
		ad.Normalizer(c)
		assert.Equal(t, "M", c.Ext[untdid.ExtKeyTaxCategory].String())
		assert.Equal(t, "7%", c.Percent.String())
	})

	t.Run("missing rate", func(t *testing.T) {
		c := &tax.Combo{
			Category: tax.CategoryVAT,
		}
		ad.Normalizer(c)
		assert.Empty(t, c.Ext)
	})
}

func TestTaxComboValidation(t *testing.T) {
	ad := tax.AddonForKey(en16931.V2017)
	t.Run("standard VAT rate", func(t *testing.T) {
		c := &tax.Combo{
			Category: tax.CategoryVAT,
			Rate:     tax.RateStandard,
			Percent:  num.NewPercentage(19, 2),
		}
		ad.Normalizer(c)
		assert.NoError(t, ad.Validator(c))
		assert.Equal(t, "S", c.Ext[untdid.ExtKeyTaxCategory].String())
		assert.Equal(t, "19%", c.Percent.String())
	})

	t.Run("exempt reverse charge", func(t *testing.T) {
		c := &tax.Combo{
			Category: tax.CategoryVAT,
			Rate:     tax.RateExempt.With(tax.TagReverseCharge),
		}
		ad.Normalizer(c)
		assert.NoError(t, ad.Validator(c))
		assert.Equal(t, "AE", c.Ext[untdid.ExtKeyTaxCategory].String())
		assert.Nil(t, c.Percent)
	})

	t.Run("nil", func(t *testing.T) {
		var tc *tax.Combo
		err := ad.Validator(tc)
		assert.NoError(t, err)
	})

	t.Run("missing rate", func(t *testing.T) {
		c := &tax.Combo{
			Category: tax.CategoryVAT,
			Percent:  num.NewPercentage(19, 2),
		}
		ad.Normalizer(c)
		err := ad.Validator(c)
		assert.ErrorContains(t, err, "ext: (untdid-tax-category: required.)")
	})

}
