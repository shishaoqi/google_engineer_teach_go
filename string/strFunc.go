package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	str := "hello world"

	// 前缀与后缀
	subStr := "he"
	prefix := strings.HasPrefix(str, subStr)
	fmt.Printf("T/F Does the string \"%s\" have prefix %s\n", str, subStr)
	fmt.Println(prefix)

	subStr = "world"
	suffix := strings.HasSuffix(str, subStr)
	fmt.Printf("T/F Does the string \"%s\" have suffix %s\n", str, subStr)
	fmt.Println(suffix)

	// 字符串包含关系
	subStr = "lo"
	contains := strings.Contains(str, subStr)
	fmt.Printf("T/F Does the string \"%s\" contain substring %s\n", str, subStr)
	fmt.Println(contains)

	// 判断子字符串或字符在父字符串中出现的位置(索引)
	subStr = "lo"
	index := strings.Index(str, subStr)
	subStr2 := "l"
	lastIndex := strings.LastIndex(str, subStr)
	fmt.Printf("%s first appear in index %d , %s last appear on index %d\n", subStr, index, subStr2, lastIndex)
	//如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位： strings.IndexRune(s string, ch int) int

	// 字符串替换
	// strings.Replace(str, new, old) string

	// 统计字符串出现的次数
	//strings.Count(s, str string) int

	//重复字符串
	//strings.Repeat(s, count int) stringß

	//修改字符大小写
	//strings.ToLower(s) string
	//strings.ToUpper(s) string

	//修剪字符串
	//strings.TrimSpace(s) string
	// Trim   TrimLeft   TrimRight

	//分割字符串
	//strings.Fields(s) []string

	//拼接字符串
	//strings.Join(a []string, sep string)
	str = "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	//从字符串读取内容
	/*
		type Reader struct {
			s	string // 对应的字符串
			i	int64  // 当前读取到的位置
			prevRune int
		}
	 */
	r := strings.NewReader("abcdefghijklmn")
	buf := make([]byte, 5)
	//Read -- Seek
	readLen, err := r.Read(buf)
	fmt.Printf("Reader %#v \n", r)
	fmt.Println("读取到的长度：", readLen)
	fmt.Println("剩余未读的长度：", r.Len())
	fmt.Println("已读取的内容：", string(buf))
	if err != nil {
		fmt.Println("错误：", err)
	}
	//ReadAt
	bufAt := make([]byte, 256)
	r.ReadAt(bufAt, 10)
	fmt.Println(string(bufAt))
	fmt.Println("剩余未读的长度：", r.Len())
	fmt.Println("已读取的内容：", string(buf))
	// ReadByte, UnreadByte
	// Seek


	//字符串与其它类型的转换
	val, _ := strconv.Atoi("123")
	fmt.Println("assci to int:", val)
	fmt.Println("int to assci:", strconv.Itoa(123))

	f := 9.01
	flostStr := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println("strconv formatFloat:", flostStr)
	float, _ := strconv.ParseFloat("800.101", 32)
	fmt.Printf("parseFloat:%f, %s \n", float, reflect.TypeOf(float))


}
