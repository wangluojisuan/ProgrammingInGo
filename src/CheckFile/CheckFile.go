package main

import (
    "fmt"
    "github.com/xuri/excelize"
)

func main() {
    xlsx, err := excelize.OpenFile("test.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)

    rows := xlsx.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}