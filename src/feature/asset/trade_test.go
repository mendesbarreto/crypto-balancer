package asset

import (
	"strings"
	"testing"
)

func TestCanTrade(test *testing.T) {
	account := setup()

	validAssets := []string{"BTC", "ETH"}

	err := CanTrade(account, validAssets)

	if err != nil {
		test.Error(err)
	}

	validAssets = []string{"BTCC", "ETH"}
	err = CanTrade(account, validAssets)

	if err == nil {
		test.Error("Should raise a error, when a asset is invalid")
	}

	validAssets = []string{"LTC", "ETH"}
	err = CanTrade(account, validAssets)
	expectedMessage := "[APIError] code=-1, msg=LTC does not have amount available for trade"

	if err != nil {
		if strings.Compare(err.Error(), expectedMessage) != 0 {
			test.Errorf("The error mesage should be %s", expectedMessage)
		}
	} else {
		test.Error("First asset on the list does not have a valid amout")
	}

	validAssets = []string{"ETH", "LTC"}
	err = CanTrade(account, validAssets)

	if err != nil {
		if strings.Compare(err.Error(), expectedMessage) != 0 {
			test.Errorf("The error mesage should be %s", expectedMessage)
		}
	} else {
		test.Error("Second asset on the list does not have a valid amout")
	}

}
