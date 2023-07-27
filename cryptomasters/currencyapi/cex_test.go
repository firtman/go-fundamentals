package currencyapi_test

import (
	"testing"

	"frontendmasters.com/cryptomasters/currencyapi"
)

func TestGetRate(t *testing.T) {
	_, err := currencyapi.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}
