package Fingerprint

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
)

func baseId() string {
	type win32BaseBoard struct {
		Model        string
		Manufacturer string
		Name         string
		SerialNumber string
	}
	var dst []win32BaseBoard
	if err := wmi.Query("SELECT * FROM Win32_BaseBoard", &dst); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s%s%s%s", dst[0].Manufacturer, dst[0].Model, dst[0].Name, dst[0].SerialNumber)
}
