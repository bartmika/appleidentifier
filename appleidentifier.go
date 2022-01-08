package appleidentifier

// GetProductNameByDeviceCode will lookup the device code (e.g. `iPhone1,1`,
// `Watch1,1`, etc.) and return the product name of the Apple device.
func GetProductNameByDeviceCode(deviceCode string) string {
	return codes[deviceCode]
}
