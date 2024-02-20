/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/9 16:11:14 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/consts/common/time.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commonconsts

var (
	WeekMap      map[int]string
	ChineseMonth map[int]string
	ChineseDay   map[int]string
	TwelveTime   map[int]string
)

func init() {
	WeekMap = map[int]string{
		1: "星期一",
		2: "星期二",
		3: "星期三",
		4: "星期四",
		5: "星期五",
		6: "星期六",
		7: "星期日",
	}
	ChineseMonth = map[int]string{
		1:  "正月",
		2:  "二月",
		3:  "三月",
		4:  "四月",
		5:  "五月",
		6:  "六月",
		7:  "七月",
		8:  "八月",
		9:  "九月",
		10: "十月",
		11: "十一月",
		12: "腊月",
	}
	ChineseDay = map[int]string{
		1:  "初一",
		2:  "初二",
		3:  "初三",
		4:  "初四",
		5:  "初五",
		6:  "初六",
		7:  "初七",
		8:  "初八",
		9:  "初九",
		10: "初十",
		11: "十一",
		12: "十二",
		13: "十三",
		14: "十四",
		15: "十五",
		16: "十六",
		17: "十七",
		18: "十八",
		19: "十九",
		20: "二十",
		21: "廿一",
		22: "廿二",
		23: "廿三",
		24: "廿四",
		25: "廿五",
		26: "廿六",
		27: "廿七",
		28: "廿八",
		29: "廿九",
		30: "三十",
	}
	TwelveTime = map[int]string{
		0:  "半夜 12点",
		1:  "凌晨 1点",
		2:  "凌晨 2点",
		3:  "凌晨 3点",
		4:  "凌晨 4点",
		5:  "凌晨 5点",
		6:  "早上 6点",
		7:  "早上 7点",
		8:  "早上 8点",
		9:  "上午 9点",
		10: "上午 10点",
		11: "中午 11点",
		12: "中午 12点",
		13: "下午 1点",
		14: "下午 2点",
		15: "下午 3点",
		16: "下午 4点",
		17: "傍晚 5点",
		18: "晚上 6点",
		19: "晚上 7点",
		20: "晚上 8点",
		21: "晚上 9点",
		22: "晚上 10点",
		23: "晚上 11点",
	}
}
