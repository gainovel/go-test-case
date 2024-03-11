# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 10:54:47 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/chan.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# features_usages_001_test.go
001/all: 001/chan_crud
001/chan_crud:
	go test -v -run TestName_2024_01_09_14_59_43/chan_crud github.com/gainovel/testcase/stdlib/runtime/chan
001/1.read_no_buf:
	go test -v -run TestName_2024_01_09_14_59_43/1.read_no_buf github.com/gainovel/testcase/stdlib/runtime/chan
001/2.read_no_data:
	go test -v -run TestName_2024_01_09_14_59_43/2.read_no_data github.com/gainovel/testcase/stdlib/runtime/chan
001/3.read_nil_chan:
	go test -v -run TestName_2024_01_09_14_59_43/3.read_nil_chan github.com/gainovel/testcase/stdlib/runtime/chan
001/1.write_no_buf:
	go test -v -run TestName_2024_01_09_14_59_43/1.write_no_buf github.com/gainovel/testcase/stdlib/runtime/chan
001/2.write_full_data:
	go test -v -run TestName_2024_01_09_14_59_43/2.write_full_data github.com/gainovel/testcase/stdlib/runtime/chan
001/3.write_nil_chan:
	go test -v -run TestName_2024_01_09_14_59_43/3.write_nil_chan github.com/gainovel/testcase/stdlib/runtime/chan
001/1.write_to_a_closed_chan:
	go test -v -run TestName_2024_01_09_14_59_43/1.write_to_a_closed_chan github.com/gainovel/testcase/stdlib/runtime/chan
001/2.close_a_closed_chan:
	go test -v -run TestName_2024_01_09_14_59_43/2.close_a_closed_chan github.com/gainovel/testcase/stdlib/runtime/chan
001/read_from_a_close_chan:
	go test -v -run TestName_2024_01_09_14_59_43/read_from_a_close_chan github.com/gainovel/testcase/stdlib/runtime/chan
