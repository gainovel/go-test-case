/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/5 14:39:32 星期五
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/tools/common/file.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package commontools

import (
	"io"
	"log"
	"os"
)

func CopyFile(dst, src string) {
	fo, err := os.Open(src)
	if err != nil {
		log.Printf("os.Open(%s) failed, err: %v", src, err)
		return
	}
	fw, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("os.OpenFile(%s) failed, err: %v", dst, err)
		return
	}
	defer fw.Close()
	defer fo.Close()
	_, err = io.Copy(fw, fo)
	if err != nil {
		log.Printf("io.Copy(%s,%s) failed, err: %v", src, dst, err)
		return
	}
}
