package asset

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetBalance(test *testing.T) {
	account := setup()

	expectedAsset := "ETH"
	balance, err := GetBalance(account.Balances, "ETH")

	if err != nil {
		test.Errorf(err.Error())
	} else {
		if balance == nil {
			test.Errorf("Asset should not be nil")
		}

		if strings.Compare(balance.Asset, expectedAsset) != 0 {
			test.Errorf("The returned asset should be the same as passed at method parameter")
		}
	}

	invalidAsset := "INVALID"
	balance, err = GetBalance(account.Balances, "INVALID")

	if err == nil {
		test.Errorf("Hen a invalid asset is passed it should raise an error")
	} else {
		expectedMessage := fmt.Sprintf("%s seems to be invalid, %s was not found in this account", invalidAsset, invalidAsset)
		if strings.Compare(err.Message, expectedMessage) != 0 {
			test.Errorf("The error message should be the same")
		}
	}

	invalidAsset = "BTC"
	duplicatedList := append(account.Balances, account.Balances[0])

	balance, err = GetBalance(duplicatedList, invalidAsset)

	if err == nil {
		test.Errorf("Hen a invalid asset is passed it should raise an error")
	} else {
		expectedMessage := fmt.Sprintf("Asset %s has more then one entry in the list", invalidAsset)
		if strings.Compare(err.Message, expectedMessage) != 0 {
			test.Errorf("The error message should \n %s \n got \n %s", err.Message, expectedMessage)
		}
	}
}
