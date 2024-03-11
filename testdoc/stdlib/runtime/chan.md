```text
/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-11 10:55:38 星期一
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : testdoc/stdlib/runtime/chan.md
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/
```

# Makefile test script

## features_usages_001_test.go


```shell
make 001/chan_crud -f Makefiles/stdlib/runtime/chan.mk
make 001/1.read_no_buf -f Makefiles/stdlib/runtime/chan.mk
make 001/2.read_no_data -f Makefiles/stdlib/runtime/chan.mk
make 001/3.read_nil_chan -f Makefiles/stdlib/runtime/chan.mk
make 001/1.write_no_buf -f Makefiles/stdlib/runtime/chan.mk
make 001/2.write_full_data -f Makefiles/stdlib/runtime/chan.mk
make 001/3.write_nil_chan -f Makefiles/stdlib/runtime/chan.mk
make 001/1.write_to_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
make 001/2.close_a_closed_chan -f Makefiles/stdlib/runtime/chan.mk
make 001/read_from_a_close_chan -f Makefiles/stdlib/runtime/chan.mk
```


