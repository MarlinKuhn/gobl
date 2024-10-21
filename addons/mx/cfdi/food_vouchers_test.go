package cfdi_test

import (
	"testing"

	"github.com/invopop/gobl/addons/mx/cfdi"
	"github.com/invopop/gobl/num"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvalidFoodVouchers(t *testing.T) {
	fvc := &cfdi.FoodVouchers{}

	err := fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "account_number: cannot be blank")
	assert.Contains(t, err.Error(), "lines: cannot be blank")

	fvc.EmployerRegistration = "123456789012345678901"
	fvc.AccountNumber = "012345678901234567891"

	err = fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "employer_registration: the length must be no more than 20")
	assert.Contains(t, err.Error(), "account_number: the length must be no more than 20")
}

func TestInvalidFoodVouchersLine(t *testing.T) {
	fvc := &cfdi.FoodVouchers{Lines: []*cfdi.FoodVouchersLine{{}}}

	err := fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "e_wallet_id: cannot be blank")
	assert.Contains(t, err.Error(), "issue_date_time: required")
	assert.Contains(t, err.Error(), "employee: cannot be blank")

	fvc.Lines[0].EWalletID = "123456789012345678901"

	err = fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "e_wallet_id: the length must be no more than 20")
}

func TestInvalidFoodVouchersEmployee(t *testing.T) {
	fvc := &cfdi.FoodVouchers{Lines: []*cfdi.FoodVouchersLine{{Employee: &cfdi.FoodVouchersEmployee{}}}}

	err := fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "tax_code: cannot be blank")
	assert.Contains(t, err.Error(), "curp: cannot be blank")
	assert.Contains(t, err.Error(), "name: cannot be blank")

	fvc.Lines[0].Employee.TaxCode = "INVALID1"
	fvc.Lines[0].Employee.CURP = "INVALID2"
	fvc.Lines[0].Employee.SocialSecurity = "INVALID3"

	err = fvc.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "tax_code: invalid tax identity code")
	assert.Contains(t, err.Error(), "curp: must be in a valid format")
	assert.Contains(t, err.Error(), "social_security: must be in a valid format")
}

func TestCalculateFoodVouchers(t *testing.T) {
	fvc := &cfdi.FoodVouchers{
		Lines: []*cfdi.FoodVouchersLine{
			{Amount: num.MakeAmount(1234, 3)},
			{Amount: num.MakeAmount(4321, 3)},
		},
	}

	err := fvc.Calculate()

	require.NoError(t, err)
	assert.Equal(t, num.MakeAmount(123, 2), fvc.Lines[0].Amount)
	assert.Equal(t, num.MakeAmount(555, 2), fvc.Total)
}
