package bill

import (
	"github.com/invopop/gobl/cal"
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/tax"
	"github.com/invopop/validation"
)

// Ordering provides additional information about the ordering process including references
// to other documents and alternative parties involved in the order-to-delivery process.
type Ordering struct {
	// Identifier assigned by the customer or buyer for internal routing purposes.
	Code cbc.Code `json:"code,omitempty" jsonschema:"title=Code" en16931:"BT-10"`
	// Any additional Codes, IDs, SKUs, or other regional or custom
	// identifiers that may be used to identify the order.
	Identities []*org.Identity `json:"identities,omitempty" jsonschema:"title=Identities"`
	// Period of time that the invoice document refers to often used in addition to the details
	// provided in the individual line items.
	Period *cal.Period `json:"period,omitempty" jsonschema:"title=Period" en16931:"BG-14"`
	// Party who is responsible for making the purchase, but is not responsible
	// for handling taxes.
	Buyer *org.Party `json:"buyer,omitempty" jsonschema:"title=Buyer"`
	// Party who is selling the goods but is not responsible for taxes like the
	// supplier.
	Seller *org.Party `json:"seller,omitempty" jsonschema:"title=Seller"`
	// Projects this invoice refers to.
	Projects []*org.DocumentRef `json:"projects,omitempty" jsonschema:"title=Projects" en16931:"BT-11"`
	// The identification of contracts.
	Contracts []*org.DocumentRef `json:"contracts,omitempty" jsonschema:"title=Contracts" en16931:"BT-12"`
	// Purchase orders issued by the customer or buyer.
	Purchases []*org.DocumentRef `json:"purchases,omitempty" jsonschema:"title=Purchase Orders" en16931:"BT-13"`
	// Sales orders issued by the supplier or seller.
	Sales []*org.DocumentRef `json:"sales,omitempty" jsonschema:"title=Sales Orders" en16931:"BT-14"`
	// Receiving Advice.
	Receiving []*org.DocumentRef `json:"receiving,omitempty" jsonschema:"title=Receiving Advice" en16931:"BT-15"`
	// Despatch advice.
	Despatch []*org.DocumentRef `json:"despatch,omitempty" jsonschema:"title=Despatch Advice" en16931:"BT-16"`
	// Tender advice, the identification of the call for tender or lot the invoice relates to.
	Tender []*org.DocumentRef `json:"tender,omitempty" jsonschema:"title=Tender Advice" en16931:"BT-17"`
}

// Normalize attempts to clean and normalize the Ordering data.
func (o *Ordering) Normalize(normalizers tax.Normalizers) {
	if o == nil {
		return
	}
	o.Code = cbc.NormalizeCode(o.Code)
	normalizers.Each(o)
	tax.Normalize(normalizers, o.Identities)
	tax.Normalize(normalizers, o.Projects)
	tax.Normalize(normalizers, o.Contracts)
	tax.Normalize(normalizers, o.Purchases)
	tax.Normalize(normalizers, o.Sales)
	tax.Normalize(normalizers, o.Receiving)
	tax.Normalize(normalizers, o.Despatch)
	tax.Normalize(normalizers, o.Tender)
	tax.Normalize(normalizers, o.Buyer)
	tax.Normalize(normalizers, o.Seller)
}

// Validate the ordering details.
func (o *Ordering) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.Code),
		validation.Field(&o.Identities),
		validation.Field(&o.Projects),
		validation.Field(&o.Contracts),
		validation.Field(&o.Purchases),
		validation.Field(&o.Sales),
		validation.Field(&o.Receiving),
		validation.Field(&o.Despatch),
		validation.Field(&o.Tender),
		validation.Field(&o.Buyer),
		validation.Field(&o.Seller),
	)
}
