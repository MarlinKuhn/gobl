package fr

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/regimes/common"
	"github.com/invopop/gobl/tax"
	"github.com/invopop/validation"
)

var (
	taxCodeVATRegexp   = regexp.MustCompile(`^\d{11}$`)
	taxCodeSIRENRegexp = regexp.MustCompile(`^\d{9}$`)
)

// normalizeTaxIdentity normalizes the SIREN code, if there are any errors,
// these will be picked up by validation.
func normalizeTaxIdentity(tID *tax.Identity) {
	if tID.Code == "" {
		return
	}
	tax.NormalizeIdentity(tID)

	str := tID.Code.String()
	// Check if we have a SIREN so we can try and normalize with the
	// check digit.
	if len(str) == 9 {
		if err := validateSIRENTaxCode(tID.Code); err != nil {
			return
		}
		chk := calculateVATCheckDigit(str)
		tID.Code = cbc.Code(fmt.Sprintf("%s%s", chk, str))
	}
}

// validateTaxIdentity checks to ensure the SIRET code looks okay.
func validateTaxIdentity(tID *tax.Identity) error {
	return validation.ValidateStruct(tID,
		validation.Field(&tID.Code, validation.By(validateVATTaxCode)),
	)
}

func validateVATTaxCode(value interface{}) error {
	code, ok := value.(cbc.Code)
	if !ok || code == "" {
		return nil
	}
	str := code.String()

	if !taxCodeVATRegexp.MatchString(str) {
		return errors.New("invalid format")
	}

	// Extract the last nine digits as an integer.
	siren := str[2:] // extract last nine digits
	chk := calculateVATCheckDigit(siren)
	expectStr := str[:2] // compare with first two digits
	if chk != expectStr {
		return errors.New("checksum mismatch")
	}

	return nil
}

func calculateVATCheckDigit(str string) string {
	// Assume we have a SIREN
	total, _ := strconv.Atoi(str)
	total = (total*100 + 12) % 97

	return fmt.Sprintf("%02d", total)
}

func validateSIRENTaxCode(value interface{}) error {
	code, ok := value.(cbc.Code)
	if !ok || code == "" {
		return nil
	}
	str := code.String()

	if !taxCodeSIRENRegexp.MatchString(str) {
		return errors.New("invalid format")
	}

	base := str[:8]
	chk := str[8:]
	v := common.ComputeLuhnCheckDigit(base)
	if chk != v {
		return errors.New("checksum mismatch")
	}

	return nil
}
