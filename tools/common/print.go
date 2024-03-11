/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 19:04:16 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/print.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	"fmt"
	commonconsts "github.com/gainovel/testcase/consts/common"

	"github.com/fatih/color"
)

func Print88StarLine() {
	fmt.Println(Generate88StarLine())
}
func Print88HorizontalLine() {
	PrintNHorizontalLine(88)
}
func Print100HorizontalLine() {
	PrintNHorizontalLine(100)
}
func Print120HorizontalLine() {
	PrintNHorizontalLine(120)
}
func PrintNHorizontalLine(n int) {
	fmt.Println(GenerateHorizontalLine(n))
}
func PrintDownHand(str string) {
	fmt.Printf("   (%s)👇(%s)\n", GenerateSimpleCurFormatTime(), str)
}
func Println(IsKeyValue bool, a ...any) {
	if IsKeyValue {
		for i := 0; i < len(a); {
			if i == 2 {
				Print88HorizontalLine()
			}
			fmt.Printf("%c ", commonconsts.VERTICAL)
			color.New(color.FgHiMagenta).Printf("%-20v", a[i])
			fmt.Printf("%s", "|  ")
			color.New(color.FgHiCyan).Printf("%v\n", a[i+1])
			i += 2
		}
	} else {
		for i := 0; i < len(a); i++ {
			if i == 2 {
				Print120HorizontalLine()
			}
			color.New(color.FgHiBlue).Printf("%c %v\n", commonconsts.VERTICAL, a[i])
		}
	}
}
func printCurFormatTime() {
	fmt.Printf("%c %s\n", commonconsts.VERTICAL, GenerateCurFormatTime())
}
func PrintAll(IsKeyValue bool, DownHandStr string, a ...any) {
	PrintDownHand(DownHandStr)
	Print88HorizontalLine()
	Println(IsKeyValue, a...)
	Print88HorizontalLine()
}
