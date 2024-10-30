package bill

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/tax"
	"github.com/invopop/jsonschema"
	"github.com/invopop/validation"
)

// LineCharge represents an amount added to the line, and will be
// applied before taxes.
type LineCharge struct {
	// Key for grouping or identifying charges for tax purposes. A suggested list of
	// keys is provided, but these are for reference only and may be extended by
	// the issuer.
	Key cbc.Key `json:"key,omitempty" jsonschema:"title=Key"`
	// Reference or ID for this charge defined by the issuer
	Code cbc.Code `json:"code,omitempty" jsonschema:"title=Code"`
	// Text description as to why the charge was applied
	Reason string `json:"reason,omitempty" jsonschema:"title=Reason"`
	// Percentage if fixed amount not applied
	Percent *num.Percentage `json:"percent,omitempty" jsonschema:"title=Percent"`
	// Fixed or resulting charge amount to apply (calculated if percent present).
	Amount num.Amount `json:"amount" jsonschema:"title=Amount" jsonschema_extras:"calculated=true"`
	// Extension codes that apply to the charge
	Ext tax.Extensions `json:"ext,omitempty" jsonschema:"title=Extensions"`
}

// Normalize performs normalization on the charge and embedded objects using the
// provided list of normalizers.
func (lc *LineCharge) Normalize(normalizers tax.Normalizers) {
	lc.Code = cbc.NormalizeCode(lc.Code)
	lc.Ext = tax.CleanExtensions(lc.Ext)
	normalizers.Each(lc)
}

// Validate checks the line charge's fields.
func (lc *LineCharge) Validate() error {
	return validation.ValidateStruct(lc,
		validation.Field(&lc.Key),
		validation.Field(&lc.Code),
		validation.Field(&lc.Percent),
		validation.Field(&lc.Amount, validation.Required, num.NotZero),
		validation.Field(&lc.Ext),
	)
}

// JSONSchemaExtend adds the charge key definitions to the schema.
func (LineCharge) JSONSchemaExtend(schema *jsonschema.Schema) {
	extendJSONSchemaWithChargeKey(schema)
}
