/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/5 14:39:32 星期五
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/time.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	"fmt"
	commonconsts "github.com/gainovel/testcase/consts/common"
	"github.com/golang-module/carbon/v2"
)

// GenerateFormatTime
/**
 * @author: Gainovel
 * @description: "2024/01/05 10:53:08" 👉 "☀️︎阳历 2024年01月05日 星期五 | 🌙农历 兔年(2023)十一月廿四 | 🕛上午 10点53分08秒"
 * @date: 2024/1/5 10:52:08
 * @params: 时间字符串(e.g. "2024/01/05 10:53:08")
 * @return 格式化字符串(e.g. "☀️︎阳历 2024年01月05日 星期五 | 🌙农历 兔年(2023)十一月廿四 | 🕛上午 10点53分08秒")
 */
func GenerateFormatTime(time string) string {
	now := carbon.Parse(time)
	SolarCalendar := fmt.Sprintf("%s︎阳历 %d年%02d月%02d日 %s",
		"☀️",
		now.Year(),
		now.Month(),
		now.Day(),
		commonconsts.WeekMap[now.DayOfWeek()],
	)
	ChineseCalendar := fmt.Sprintf("%s农历 %s年(%d)%s%s",
		"🌙",
		now.Lunar().Animal(),
		now.Lunar().Year(),
		now.Lunar().ToMonthString(),
		now.Lunar().ToDayString(),
	)
	TodayTime := fmt.Sprintf("%s%s%02d分%02d秒",
		"🕛",
		commonconsts.TwelveTime[now.Hour()],
		now.Minute(),
		now.Second(),
	)
	formatTime := fmt.Sprintf("%s | %s | %s", SolarCalendar, ChineseCalendar, TodayTime)
	return formatTime
}
func GenerateSimpleFormatTime(time string) string {
	now := carbon.Parse(time)
	SolarCalendar := fmt.Sprintf("%s%d年%02d月%02d日",
		"☀️",
		now.Year(),
		now.Month(),
		now.Day(),
	)
	TodayTime := fmt.Sprintf("%s%s%02d分%02d秒",
		"🕛",
		commonconsts.TwelveTime[now.Hour()],
		now.Minute(),
		now.Second(),
	)
	formatTime := fmt.Sprintf("%s%s", SolarCalendar, TodayTime)
	return formatTime
}

// GenerateCurFormatTime
/**
 * @author: Gainovel
 * @description: 生成当前时间的格式化输出
 * @date: 2024/1/5 11:12:08
 * @params:
 * @return: string(e.g. "2024-01-05")
 */
func GenerateCurFormatTime() string {
	return GenerateFormatTime(carbon.Now().String())
}
func GenerateSimpleCurFormatTime() string {
	return GenerateSimpleFormatTime(carbon.Now().String())
}
