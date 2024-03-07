package co

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/regimes/common"
	"github.com/invopop/gobl/tax"
	"github.com/invopop/validation"
)

// Known base tax identity types for Colombia
const (
	TaxIdentityTypeTIN        cbc.Key = "tin"
	TaxIdentityTypeCard       cbc.Key = "card"
	TaxIdentityTypeCitizen    cbc.Key = "citizen"
	TaxIdentityTypePassport   cbc.Key = "passport"
	TaxIdentityTypeIndividual cbc.Key = "individual"
	TaxIdentityTypeCivil      cbc.Key = "civil"
	TaxIdentityTypeForeign    cbc.Key = "foreign"
	TaxIdentityTypeForeigner  cbc.Key = "foreigner"
	TaxIdentityTypePEP        cbc.Key = "pep"
	TaxIdentityTypeNUIP       cbc.Key = "nuip"

	TaxCodeFinalCustomer cbc.Code = "222222222222"
)

var (
	nitMultipliers = []int{3, 7, 13, 17, 19, 23, 29, 37, 41, 43, 47, 53, 59, 67, 71}
)

var taxIdentityTypeDefs = []*tax.KeyDefinition{
	{
		Key: TaxIdentityTypeTIN, // DEFAULT!
		Name: i18n.String{
			i18n.EN: "TIN - Tax Identification Number",
			i18n.ES: "NIT - Número de Identificación Tributaria",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "31",
			KeyDIANAdditionalAccountID: "1",
		},
	},
	{
		Key: TaxIdentityTypeCivil,
		Name: i18n.String{
			i18n.ES: "Registro Civil",
			i18n.EN: "Civil Registry",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "11",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeCard,
		Name: i18n.String{
			i18n.EN: "Identity Card",
			i18n.ES: "Tarjeta de Identidad",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "12",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeCitizen,
		Name: i18n.String{
			i18n.EN: "Citizen Identity Card",
			i18n.ES: "Cédula de ciudadanía",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "13",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeCard.With(TaxIdentityTypeForeigner),
		Name: i18n.String{
			i18n.EN: "Foreigner Identity Card",
			i18n.ES: "Tarjeta de Extranjería",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "21",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeCitizen.With(TaxIdentityTypeForeigner),
		Name: i18n.String{
			i18n.EN: "Foreigner Citizen Identity Card",
			i18n.ES: "Cédula de extranjería",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "22",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeTIN.With(TaxIdentityTypeIndividual),
		Name: i18n.String{
			i18n.EN: "TIN of an individual",
			i18n.ES: "NIT de persona natural",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "31",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypePassport,
		Name: i18n.String{
			i18n.EN: "Passport",
			i18n.ES: "Pasaporte",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "41",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeForeign,
		Name: i18n.String{
			i18n.EN: "Foreign Document",
			i18n.ES: "Documento de identificación extranjero",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "42",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypePEP,
		Name: i18n.String{
			i18n.EN: "PEP - Special Permit to Stay",
			i18n.ES: "PEP - Permiso Especial de Permanencia",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "47",
			KeyDIANAdditionalAccountID: "2",
		},
	},
	{
		Key: TaxIdentityTypeTIN.With(TaxIdentityTypeForeign),
		Name: i18n.String{
			i18n.EN: "Foreign TIN",
			i18n.ES: "NIT de otro país",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "50",
			KeyDIANAdditionalAccountID: "1",
		},
	},
	{
		Key: TaxIdentityTypeNUIP,
		Name: i18n.String{
			i18n.EN: "NUIP - National Unique Personal Identification Number",
			i18n.ES: "NUIP - Número Único de Identificación Personal",
		},
		Map: cbc.CodeMap{
			KeyDIANCompanyID:           "91",
			KeyDIANAdditionalAccountID: "2",
		},
	},
}

// validateTaxIdentity checks to ensure the NIT code looks okay.
func validateTaxIdentity(tID *tax.Identity) error {
	return validation.ValidateStruct(tID,
		validation.Field(&tID.Type, validation.Required),
		validation.Field(&tID.Code,
			validation.When(tID.Type.In(TaxIdentityTypeTIN),
				validation.Required,
				validation.By(validateTaxCode),
			),
		),
		validation.Field(&tID.Zone,
			validation.When(!isFinalCustomer(tID),
				validation.Required,
				tax.ZoneIn(zones),
			),
		),
	)
}

func isFinalCustomer(tID *tax.Identity) bool {
	if tID == nil {
		return false
	}
	return tID.Type == TaxIdentityTypeCitizen && tID.Code == TaxCodeFinalCustomer
}

// normalizeTaxIdentity will remove any whitespace or separation characters from
// the tax code and also make sure the default type is set.
func normalizeTaxIdentity(tID *tax.Identity) error {
	if tID == nil {
		return nil
	}
	if tID.Type == cbc.KeyEmpty {
		tID.Type = TaxIdentityTypeTIN // set the default
	}
	if err := common.NormalizeTaxIdentity(tID); err != nil {
		return err
	}
	return nil
}

func normalizePartyWithTaxIdentity(p *org.Party) error {
	// override the party's locality and region using the tax identity zone data.
	tID := p.TaxID
	if tID != nil && tID.Zone != "" {
		z := zones.Get(tID.Zone)
		if z != nil {
			if len(p.Addresses) == 0 {
				return nil
			}
			a := p.Addresses[0]
			a.Locality = z.Locality.In(i18n.ES)
			a.Region = z.Region.In(i18n.ES)
		}
	}
	return nil
}

func validateTaxCode(value interface{}) error {
	code, ok := value.(cbc.Code)
	if !ok {
		return nil
	}
	if code == "" {
		return nil
	}
	for _, v := range code {
		x := v - 48
		if x < 0 || x > 9 {
			return errors.New("contains invalid characters")
		}
	}
	l := len(code)
	if l > 10 {
		return errors.New("too long")
	}
	if l < 9 {
		return errors.New("too short")
	}

	return validateDigits(code[0:l-1], code[l-1:l])
}

func validateDigits(code, check cbc.Code) error {
	ck, err := strconv.Atoi(string(check))
	if err != nil {
		return fmt.Errorf("invalid check: %w", err)
	}

	sum := 0
	l := len(code)
	for i, v := range code {
		// 48 == ASCII "0"
		sum += int(v-48) * nitMultipliers[l-i-1]
	}
	sum = sum % 11
	if sum >= 2 {
		sum = 11 - sum
	}

	if sum != ck {
		return errors.New("checksum mismatch")
	}

	return nil
}

func taxIdentityTypeKeys() []interface{} {
	keys := make([]interface{}, len(taxIdentityTypeDefs))
	i := 0
	for _, v := range taxIdentityTypeDefs {
		keys[i] = v.Key
		i++
	}
	return keys
}

var isValidTaxIdentityTypeKey = validation.In(taxIdentityTypeKeys()...)
