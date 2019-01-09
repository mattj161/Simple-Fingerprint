package Fingerprint

import (
	"github.com/StackExchange/wmi"
	"log"
)

// Doesn't really work if run on VPS... Use at own risk.
func macId() string {
	type win32NetworkAdapterConfiguration struct {
		MACAddress string
		IPEnabled  bool
	}
	var dst []win32NetworkAdapterConfiguration
	if err := wmi.Query("SELECT * FROM Win32_NetworkAdapterConfiguration", &dst); err != nil {
		log.Fatal(err)
	}
	for _, d := range dst {
		if d.IPEnabled {
			return d.MACAddress
		}
	}
	return ""
}
