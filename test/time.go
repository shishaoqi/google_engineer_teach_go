package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main()  {
	var t string = "2021-03-25 16:24:21"

	regSn := regexp.MustCompile(`^[1-9]{1}\.\d{1,9}[E|e]\+09$`) // 匹配科学计数法表示
	regSnResult := string(regSn.Find([]byte(t)))

	if regSnResult != "" {
		occurrence_time_float64, ok := t.(float64)
		if !ok {
			fmt.Println("assertion to float64 failed")
		}
		occurrence_time, _ := RecoverScienceNotation(occurrence_time_float64)  // 科学计数法 -> 时间戳
		timestamp, _ := strconv.ParseInt(occurrence_time, 10, 64)
		tm := time.Unix(timestamp, 0)
		occurrenceTime := tm.Format("2006-01-02 15:04:05")
		fmt.Println("occurrenceTime: ", occurrenceTime)
	}
}

// 科学计数法->正常int表示
func RecoverScienceNotation(sn float64) (string ,error){
	var (
		new float64
	)

	old  := strconv.FormatFloat(sn, 'E', -1, 64)
	n, err := fmt.Sscanf(old, "%e", &new)
	if err != nil {
		fmt.Println(err)
		return "", nil
	} else if 1 != n {
		fmt.Println("n is not one")
		return "", nil
	}
	return strconv.FormatInt(int64(new), 10), nil
}