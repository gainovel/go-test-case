/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-01-14 15:39:02 星期日
 * @ProductName  : GoLand
 * @PrjectName   : go-examples
 * @File         : examples/thirdparty/github.com/github.com_gookit_color/color_quickstart_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package github_com_gookit_color

import (
	"fmt"
	"github.com/gookit/color"
	"testing"
)

func TestName_2024_01_14_15_39_02(t *testing.T) {
	// go test -run TestName_2024_01_14_15_39_02/quick_start
	t.Run("quick start", func(t *testing.T) {
		// quick use package func
		color.Redp("Simple to use color")
		color.Redln("Simple to use color")
		color.Greenp("Simple to use color\n")
		color.Cyanln("Simple to use color")
		color.Yellowln("Simple to use color")

		// quick use like fmt.Print*
		color.Red.Println("Simple to use color")
		color.Green.Print("Simple to use color\n")
		color.Cyan.Printf("Simple to use %s\n", "color")
		color.Yellow.Printf("Simple to use %s\n", "color")

		// use like func
		red := color.FgRed.Render
		green := color.FgGreen.Render
		fmt.Printf("%s line %s library\n", red("Command"), green("color"))

		// custom color
		color.New(color.FgWhite, color.BgBlack).Println("custom color style")

		// can also:
		color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

		// internal theme/style:
		color.Info.Tips("message")
		color.Info.Prompt("message")
		color.Info.Println("message")
		color.Warn.Println("message")
		color.Error.Println("message")

		// use style tag
		color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")
		// Custom label attr: Supports the use of 16 color names, 256 color values, rgb color values and hex color values
		color.Println("<fg=11aa23>he</><bg=120,35,156>llo</>, <fg=167;bg=232>wel</><fg=red>come</>")

		// apply a style tag
		color.Tag("info").Println("info style text")

		// prompt message
		color.Info.Prompt("prompt style message")
		color.Warn.Prompt("prompt style message")

		// tips message
		color.Info.Tips("tips style message")
		color.Warn.Tips("tips style message")
	})
}
