/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-05 14:46:36 星期一
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : examples/thirdparty/github.com/githubcom_PuerkitoBio_goquery/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package goquery000

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestName_2024_02_05_14_46_36(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		// Request the HTML page.
		res, err := http.Get("http://zhongguose.com/")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc.Text())
		//// Find the review items
		//doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		//	// For each item found, get the title
		//	title := s.Find("a").Text()
		//	fmt.Printf("Review %d: %s\n", i, title)
		//})
		doc.Find("#colors li").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find("div a span").Text()
			fmt.Printf("Review %d: %s\n", i, title)
		})

	})
	t.Run("local html", func(t *testing.T) {
		var (
			htmlReader io.Reader
			err        error
		)
		htmlReader, err = os.Open("1.html")
		fmt.Println(err)
		doc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc.Text())
		//// Find the review items
		//doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		//	// For each item found, get the title
		//	title := s.Find("a").Text()
		//	fmt.Printf("Review %d: %s\n", i, title)
		//})
		doc.Find("#colors li").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find("div a span").Text()
			fmt.Printf("Review %d: %s\n", i, title)
		})
	})
}
