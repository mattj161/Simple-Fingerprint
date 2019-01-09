package Fingerprint

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
)

func biosId() string {
	type win32Bios struct {
		Manufacturer       string
		SMBIOSBIOSVersion  string
		IdentificationCode string
		SerialNumber       string
		ReleaseDate        string
		Version            string
	}
	var dst []win32Bios
	if err := wmi.Query("SELECT * FROM Win32_BIOS", &dst); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s%s%s%s%s%s", dst[0].Manufacturer, dst[0].SerialNumber, dst[0].Version,
		dst[0].ReleaseDate, dst[0].SMBIOSBIOSVersion, dst[0].IdentificationCode)
}
