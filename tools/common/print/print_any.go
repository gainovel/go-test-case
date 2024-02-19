/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-14 16:09:40 星期日
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/tools/common/print/print_any.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commonprint

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/goutil/dump"

	commontools "gainovel.com/go/testcase/tools/common"
)

var (
	MyFmt = newMyFmt()
)

type myFmt struct {
	horizontal string
	dumpStd2   *dump.Dumper
}

func newMyFmt() *myFmt {
	return &myFmt{
		horizontal: commontools.GenerateHorizontalLine(88),
		dumpStd2:   dump.Std2(),
	}
}

func (mf *myFmt) DumpPrintln(a ...any) {
	if len(a) == 0 {
		fmt.Println()
		return
	}
	fmt.Println(mf.horizontal)
	for i := 0; i < len(a); i++ {
		mf.dumpStd2.Println(a[i])
	}
	fmt.Println(mf.horizontal)
}

func (mf *myFmt) ColorDescPrintln(a any) {
	if a == nil {
		fmt.Println()
	}
	color.HiCyan.Printf("%s %v\n", "👇", a)
}

func (mf *myFmt) KeyValuePrintln(a ...any) {
	if len(a) == 0 {
		fmt.Println()
		return
	}
	fmt.Println(mf.horizontal)
	fmt.Printf("|%18s | %s\n", "key", "value")
	fmt.Println(mf.horizontal)
	for i := 0; i < len(a); i = i + 2 {
		fmt.Printf("|%s | %s\n", color.HiRed.Sprintf("%18v", a[i]), color.HiBlue.Sprintf("%v", a[i+1]))
	}
	fmt.Println(mf.horizontal)
}
