package Fingerprint

import (
	"github.com/StackExchange/wmi"
	"log"
)

func videoID() string {
	type win32VideoController struct {
		DriverVersion string
		Name          string
	}
	var dst []win32VideoController
	if err := wmi.Query("SELECT * FROM Win32_VideoController", &dst); err != nil {
		log.Fatal(err)
	}
	return ""
}
