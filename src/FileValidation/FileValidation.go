package main

import (
    "fmt"
    "github.com/xuri/excelize"
)

func main() {
	fileInfo, _ := os.Stat("1.xlsx")
    xlsx, err := excelize.OpenFile("1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    cell := xlsx.GetCellValue("ATS系统发布项信息表", "B2")
    fmt.Println(cell)

    rows := xlsx.GetRows("ATS系统发布项信息表")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
	
	fmt.Scan()
}