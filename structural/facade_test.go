package structural

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	accountID := "account"
	securityCode := 1234
	wf := NewWalletFacade(accountID, securityCode)

	creditCases := []struct {
		message      string
		expected     error
		accountID    string
		securityCode int
		amount       int
		balance      int
	}{
		{
			"invalid account ID",
			fmt.Errorf("%s doesn't match", "account name"),
			"invalid account name",
			securityCode,
			1000,
			0,
		},
		{
			"invalid security code",
			fmt.Errorf("%s doesn't match", "security code"),
			accountID,
			111111,
			1000,
			0,
		},
		{
			"wallet balance increases successfully",
			nil,
			accountID,
			securityCode,
			1000,
			1000,
		},
	}
	for _, v := range creditCases {
		err := wf.AddMoney(v.accountID, v.securityCode, v.amount)
		if err != nil {
			assert.EqualError(t, err, v.expected.Error())
		}
		assert.Equal(t, v.balance, wf.wallet.balance)
	}

	debitCases := []struct {
		message      string
		expected     error
		accountID    string
		securityCode int
		amount       int
		balance      int
	}{
		{
			"invalid account ID",
			fmt.Errorf("%s doesn't match", "account name"),
			"invalid account name",
			securityCode,
			1000,
			1000,
		},
		{
			"invalid security code",
			fmt.Errorf("%s doesn't match", "security code"),
			accountID,
			111111,
			1000,
			1000,
		},
		{
			"insufficient wallet money",
			fmt.Errorf("not enough balance"),
			accountID,
			securityCode,
			1200,
			1000,
		},
		{
			"wallet balance increases successfully",
			nil,
			accountID,
			securityCode,
			1000,
			0,
		},
	}
	for _, v := range debitCases {
		err := wf.DeductMoney(v.accountID, v.securityCode, v.amount)
		if err != nil {
			assert.EqualError(t, err, v.expected.Error())
		}
		assert.Equal(t, v.balance, wf.wallet.balance)
	}

}
