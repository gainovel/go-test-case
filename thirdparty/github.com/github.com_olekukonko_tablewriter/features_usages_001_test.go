/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-09 12:10:58 星期六
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package tablewriter000

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"testing"
)

// TestName_2024_03_09_12_10_58
/**
 * @author: Administrator
 * @description: 三方库 github.com/olekukonko/tablewriter用法测试，测试命令已贴在相应位置，在terminal打开项目根目录执行即可
 * @date: 2024-03-09 12:28:38
 */
// make all -f Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001.mk
func TestName_2024_03_09_12_10_58(t *testing.T) {
	// make case1 -f Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001.mk
	t.Run("case1", func(t *testing.T) {
		data := [][]string{
			[]string{"A", "The Good", "500"},
			[]string{"B", "The Very very Bad Man", "288"},
			[]string{"C", "The Ugly", "120"},
			[]string{"D", "The Gopher", "800"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Sign", "Rating"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output
	})
	// make case2 -f Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001.mk
	t.Run("case2", func(t *testing.T) {
		data := [][]string{
			[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
			[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
			[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
			[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
		table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
		table.SetBorder(false)                                // Set Border to false
		table.AppendBulk(data)                                // Add Bulk Data
		table.Render()
	})
	// make case3 -f Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001.mk
	t.Run("case3", func(t *testing.T) {
		data := [][]string{
			[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
			[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
			[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
			[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	})
	// make case4 -f Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/features_usages_001.mk
	t.Run("case4", func(t *testing.T) {
		data := [][]string{
			[]string{"Test1Merge", "HelloCol2 - 1", "HelloCol3 - 1", "HelloCol4 - 1"},
			[]string{"Test1Merge", "HelloCol2 - 2", "HelloCol3 - 2", "HelloCol4 - 2"},
			[]string{"Test1Merge", "HelloCol2 - 3", "HelloCol3 - 3", "HelloCol4 - 3"},
			[]string{"Test2Merge", "HelloCol2 - 4", "HelloCol3 - 4", "HelloCol4 - 4"},
			[]string{"Test2Merge", "HelloCol2 - 5", "HelloCol3 - 5", "HelloCol4 - 5"},
			[]string{"Test2Merge", "HelloCol2 - 6", "HelloCol3 - 6", "HelloCol4 - 6"},
			[]string{"Test2Merge", "HelloCol2 - 7", "HelloCol3 - 7", "HelloCol4 - 7"},
			[]string{"Test3Merge", "HelloCol2 - 8", "HelloCol3 - 8", "HelloCol4 - 8"},
			[]string{"Test3Merge", "HelloCol2 - 9", "HelloCol3 - 9", "HelloCol4 - 9"},
			[]string{"Test3Merge", "HelloCol2 - 10", "HelloCol3 -10", "HelloCol4 - 10"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Col1", "Col2", "Col3", "Col4"})
		table.SetFooter([]string{"", "", "Footer3", "Footer4"})
		table.SetBorder(false)

		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
			tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

		table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

		table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
			tablewriter.Colors{tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiRedColor})

		colorData1 := []string{"TestCOLOR1Merge", "HelloCol2 - COLOR1", "HelloCol3 - COLOR1", "HelloCol4 - COLOR1"}
		colorData2 := []string{"TestCOLOR2Merge", "HelloCol2 - COLOR2", "HelloCol3 - COLOR2", "HelloCol4 - COLOR2"}

		for i, row := range data {
			if i == 4 {
				table.Rich(colorData1, []tablewriter.Colors{tablewriter.Colors{}, tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor}, tablewriter.Colors{}})
				table.Rich(colorData2, []tablewriter.Colors{tablewriter.Colors{tablewriter.Normal, tablewriter.FgMagentaColor}, tablewriter.Colors{}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor}, tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Italic, tablewriter.BgHiCyanColor}})
			}
			table.Append(row)
		}

		table.SetAutoMergeCells(true)
		table.Render()

	})
}
