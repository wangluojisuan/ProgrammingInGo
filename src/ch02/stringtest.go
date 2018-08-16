package main

import "fmt"

func main() {
	var chineseStr string = "一二三四五六123456"
	var englishStr string = "abcdef"
	
	fmt.Printf("chineseLen=%d\nenglishLen=%d", len([]rune(chineseStr)), len(englishStr))
}