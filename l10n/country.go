package l10n

import (
	"github.com/invopop/jsonschema"
	"github.com/invopop/validation"
)

// CountryCode defines an ISO 3166-2 country code.
type CountryCode Code

// CountryDef provides the structure use to define a Country Code
// definition.
type CountryDef struct {
	// ISO 3166-2 Country code
	Code CountryCode `json:"code" jsonschema:"ISO 3166-2 Country Code"`
	// ISO 3166-1 alpha-3 Country code
	Alpha3 string `json:"alpha3" jsonschema:"ISO 3166-1 Alpha-3 Country Code"`
	// English name of the country
	Name string `json:"name" jsonschema:"Name"`
	// Internet Top-Level-Domain
	TLD string `json:"tld" jsonschema:"Top level domain"`
}

func validCountryCodes() []interface{} {
	list := make([]interface{}, len(CountryDefinitions))
	for i, v := range CountryDefinitions {
		list[i] = string(v.Code)
	}
	return list
}

var (
	isCountry = validation.In(validCountryCodes()...)
)

// Validate ensures the country code is inside the known and valid
// list of countries.
func (c CountryCode) Validate() error {
	return validation.Validate(string(c), isCountry)
}

// JSONSchema provides a representation of the struct for usage in Schema.
func (CountryCode) JSONSchema() *jsonschema.Schema {
	s := &jsonschema.Schema{
		Title:       "Country Code",
		Type:        "string",
		OneOf:       make([]*jsonschema.Schema, len(CountryDefinitions)),
		Description: "Defines an ISO 3166-2 country code with exception cases for tax based extensions.",
	}
	for i, v := range CountryDefinitions {
		s.OneOf[i] = &jsonschema.Schema{
			Const: v.Code,
			Title: v.Name,
		}
	}
	return s
}

// In returns true if the country code is contained inside the provided set
func (c CountryCode) In(set ...CountryCode) bool {
	for _, x := range set {
		if c == x {
			return true
		}
	}
	return false
}

// String provides string representation of the country code
func (c CountryCode) String() string {
	return string(c)
}

// Name provides the Country Name for the code
func (c CountryCode) Name() string {
	for _, v := range CountryDefinitions {
		if v.Code == c {
			return v.Name
		}
	}
	return ""
}

// Alpha3 provides the ISO 3166-1 alpha-3 country code
func (c CountryCode) Alpha3() string {
	for _, v := range CountryDefinitions {
		if v.Code == c {
			return v.Alpha3
		}
	}
	return ""
}
