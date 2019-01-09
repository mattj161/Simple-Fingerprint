package Fingerprint

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
)

func diskId() string {
	type win32DiskDrive struct {
		Model        string
		Manufacturer string
		Signature    int
		TotalHeads   int
	}
	var dst []win32DiskDrive
	if err := wmi.Query("SELECT * FROM Win32_DiskDrive", &dst); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s%s%d%d", dst[0].Manufacturer, dst[0].Model, dst[0].Signature, dst[0].TotalHeads)
}
