package appleidentifier

import "testing"

func TestGetProductNameByDeviceCode(t *testing.T) {
	productName := GetProductNameByDeviceCode("iPhone1,1")
	if productName != "iPhone" {
		t.Errorf("ProductName was incorrect, got: %s, want: %s.", productName, "iPhone")
	}
}
