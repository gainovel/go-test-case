# @Author       : Gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-02-21 11:40:08 星期三
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : /Makefile
# @Version      : v0.1.0
# @Description  : 开发中···

# stdlib/runtime/chan/features_usages_001_test.go
stdlib/runtime/chan/features_usages_001/all:
	@go test -v -run TestName_2024_01_09_14_59_43/chan_crud github.com/gainovel/testcase/stdlib/runtime/chan
	@go test -v -run TestName_2024_01_09_14_59_43/read_from_a_close_chan github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/chan_crud:
	@go test -v -run TestName_2024_01_09_14_59_43/chan_crud github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/1.read_no_buf:
	@go test -v -run TestName_2024_01_09_14_59_43/1.read_no_buf github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/2.read_no_data:
	@go test -v -run TestName_2024_01_09_14_59_43/2.read_no_data github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/3.read_nil_chan:
	@go test -v -run TestName_2024_01_09_14_59_43/3.read_nil_chan github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/1.write_no_buf:
	@go test -v -run TestName_2024_01_09_14_59_43/1.write_no_buf github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/2.write_full_data:
	@go test -v -run TestName_2024_01_09_14_59_43/2.write_full_data github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/3.write_nil_chan:
	@go test -v -run TestName_2024_01_09_14_59_43/3.write_nil_chan github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/1.write_to_a_closed_chan:
	@go test -v -run TestName_2024_01_09_14_59_43/1.write_to_a_closed_chan github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/2.close_a_closed_chan:
	@go test -v -run TestName_2024_01_09_14_59_43/2.close_a_closed_chan github.com/gainovel/testcase/stdlib/runtime/chan
stdlib/runtime/chan/features_usages_001/read_from_a_close_chan:
	@go test -v -run TestName_2024_01_09_14_59_43/read_from_a_close_chan github.com/gainovel/testcase/stdlib/runtime/chan

# stdlib/runtime/slice/features_usages_001_test.go
# 命令开头加-，忽略错误继续执行Makefile
stdlib/runtime/slice/features_usages_001/all:
	-go test -v -run TestName_2024_01_09_18_10_28/index_out_of_range github.com/gainovel/testcase/stdlib/runtime/slice
	go test -v -run TestName_2024_01_09_18_10_28/share_array github.com/gainovel/testcase/stdlib/runtime/slice
	go test -v -run TestName_2024_01_09_18_10_28/slice_copy github.com/gainovel/testcase/stdlib/runtime/slice
	go test -v -run TestName_2024_01_09_18_10_28/expansion_experiment github.com/gainovel/testcase/stdlib/runtime/slice
stdlib/runtime/slice/features_usages_001/index_out_of_range:
	go test -v -run TestName_2024_01_09_18_10_28/index_out_of_range github.com/gainovel/testcase/stdlib/runtime/slice
stdlib/runtime/slice/features_usages_001/share_array:
	go test -v -run TestName_2024_01_09_18_10_28/share_array github.com/gainovel/testcase/stdlib/runtime/slice
stdlib/runtime/slice/features_usages_001/slice_copy:
	go test -v -run TestName_2024_01_09_18_10_28/slice_copy github.com/gainovel/testcase/stdlib/runtime/slice
stdlib/runtime/slice/features_usages_001/expansion_experiment:
	go test -v -run TestName_2024_01_09_18_10_28/expansion_experiment github.com/gainovel/testcase/stdlib/runtime/slice


# stdlib/runtime/map/features_usages_001_test.go
# 命令开头加-，忽略错误继续执行Makefile
stdlib/runtime/map/features_usages_001/all:
	-go test -v -run TestName_2024_01_10_11_22_46/concurrent_map_write github.com/gainovel/testcase/stdlib/runtime/map
	-go test -v -run TestName_2024_01_10_11_22_46/assignment_to_entry_in_nil_map github.com/gainovel/testcase/stdlib/runtime/map
	-go test -v -run TestName_2024_01_10_11_22_46/map_crud github.com/gainovel/testcase/stdlib/runtime/map
stdlib/runtime/map/features_usages_001/concurrent_map_write:
	-go test -v -run TestName_2024_01_10_11_22_46/concurrent_map_write github.com/gainovel/testcase/stdlib/runtime/map
stdlib/runtime/map/features_usages_001/assignment_to_entry_in_nil_map:
	-go test -v -run TestName_2024_01_10_11_22_46/assignment_to_entry_in_nil_map github.com/gainovel/testcase/stdlib/runtime/map
stdlib/runtime/map/features_usages_001/map_crud:
	-go test -v -run TestName_2024_01_10_11_22_46/map_crud github.com/gainovel/testcase/stdlib/runtime/map

# stdlib/runtime/iota/features_usages_001_test.go
# 命令开头加-，忽略错误继续执行Makefile
stdlib/runtime/iota/features_usages_001/all:
	go test -v -run TestName_2024_01_10_16_36_44/iota_1 github.com/gainovel/testcase/stdlib/runtime/iota
	go test -v -run TestName_2024_01_10_16_36_44/iota_2 github.com/gainovel/testcase/stdlib/runtime/iota
	go test -v -run TestName_2024_01_10_16_36_44/iota_3 github.com/gainovel/testcase/stdlib/runtime/iota
stdlib/runtime/iota/features_usages_001/iota_1:
	go test -v -run TestName_2024_01_10_16_36_44/iota_1 github.com/gainovel/testcase/stdlib/runtime/iota
stdlib/runtime/iota/features_usages_001/iota_2:
	go test -v -run TestName_2024_01_10_16_36_44/iota_2 github.com/gainovel/testcase/stdlib/runtime/iota
stdlib/runtime/iota/features_usages_001/iota_3:
	go test -v -run TestName_2024_01_10_16_36_44/iota_3 github.com/gainovel/testcase/stdlib/runtime/iota

# stdlib/runtime/string/features_usages_001_test.go
stdlib/runtime/string/features_usages_001/all:
	go test -v -run TestName_2024_01_10_17_38_02/string_to_bytes_or_runes github.com/gainovel/testcase/stdlib/runtime/string
stdlib/runtime/string/features_usages_001/string_to_bytes_or_runes:
	go test -v -run TestName_2024_01_10_17_38_02/string_to_bytes_or_runes github.com/gainovel/testcase/stdlib/runtime/string

# stdlib/runtime/for/for_001_test.go
stdlib/runtime/for/for_001/all:
	go test -v -run TestName_2024_01_15_11_16_48/case1 github.com/gainovel/testcase/stdlib/runtime/for
stdlib/runtime/string/features_usages_001/case1:
	go test -v -run TestName_2024_01_15_11_16_48/case1 github.com/gainovel/testcase/stdlib/runtime/for

# stdlib/struct/features_usages_001_test.go
stdlib/struct/features_usages_001/all:
	go test -v -run TestName_2024_02_20_12_18_40/case1 github.com/gainovel/testcase/stdlib/struct
stdlib/struct/features_usages_001/case1:
	go test -v -run TestName_2024_02_20_12_18_40/case1 github.com/gainovel/testcase/stdlib/struct

# stdlib/sync/map/features_usages_001_test.go
stdlib/sync/map/features_usages_001/all:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map_1 github.com/gainovel/testcase/stdlib/sync/map
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadOrStore github.com/gainovel/testcase/stdlib/sync/map
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadAndDelete github.com/gainovel/testcase/stdlib/sync/map
stdlib/sync/map/features_usages_001/sync.Map_1:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map_1 github.com/gainovel/testcase/stdlib/sync/map
stdlib/sync/map/features_usages_001/sync.Map.LoadOrStore:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadOrStore github.com/gainovel/testcase/stdlib/sync/map
stdlib/sync/map/features_usages_001/sync.Map.LoadAndDelete:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadAndDelete github.com/gainovel/testcase/stdlib/sync/map