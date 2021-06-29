package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	newReader := transform.NewReader(resp.Body, e.NewDecoder()) // simplifiedchinese.GBK.NewDecoder()
	body, err := ioutil.ReadAll(newReader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", body)
	printCityList(body)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-1573aa7c>([^<]+)</a>`)
	submatch := compile.FindAllSubmatch(contents, -1)
	count := 0
	for _, m := range submatch {
		fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
		count++
	}
	fmt.Printf("City count: %v", count)
}
