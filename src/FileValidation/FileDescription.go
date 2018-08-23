package main

import (
    "fmt"
    "github.com/gonutz/w32"
)

func main() {
    //const path = `CUIT.exe`
	const path = `log4net.dll`

    size := w32.GetFileVersionInfoSize(path)
    if size <= 0 {
        panic("GetFileVersionInfoSize failed")
    }

    info := make([]byte, size)
    ok := w32.GetFileVersionInfo(path, info)
    if !ok {
        panic("GetFileVersionInfo failed")
    }

    fixed, ok := w32.VerQueryValueRoot(info)
    if !ok {
        panic("VerQueryValueRoot failed")
    }
    version := fixed.FileVersion()
    fmt.Printf(
        "file version: %d.%d.%d.%d\n",
        version&0xFFFF000000000000>>48,
        version&0x0000FFFF00000000>>32,
        version&0x00000000FFFF0000>>16,
        version&0x000000000000FFFF>>0,
    )

    translations, ok := w32.VerQueryValueTranslations(info)
    if !ok {
        panic("VerQueryValueTranslations failed")
    }
    if len(translations) == 0 {
        panic("no translation found")
    }
    fmt.Println("translations:", translations)

    t := translations[0]
    // w32.CompanyName simply translates to "CompanyName"
    company, ok := w32.VerQueryValueString(info, t, w32.CompanyName)
    if !ok {
        panic("cannot get company name")
    }
    fmt.Println("company:", company)
}