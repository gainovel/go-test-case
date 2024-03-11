# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-09 12:18:32 星期六
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/thirdparty/github.com/github.com_olekukonko_tablewriter/tablewriter_001.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: case1 case2 case3 case4

case1:
	go test -v -run TestName_2024_03_09_12_10_58/case1 github.com/gainovel/testcase/thirdparty/github.com/github.com_olekukonko_tablewriter
case2:
	go test -v -run TestName_2024_03_09_12_10_58/case2 github.com/gainovel/testcase/thirdparty/github.com/github.com_olekukonko_tablewriter
case3:
	go test -v -run TestName_2024_03_09_12_10_58/case3 github.com/gainovel/testcase/thirdparty/github.com/github.com_olekukonko_tablewriter
case4:
	go test -v -run TestName_2024_03_09_12_10_58/case4 github.com/gainovel/testcase/thirdparty/github.com/github.com_olekukonko_tablewriter

