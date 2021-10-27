package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"com.sand.airdroidkids":{"h":1,"count":1,"time":1261,"flow_wifi":302707,"flow_mobile":0,"time_stamp":1630285200000},"master.app.screentime":{"notification":"2"},"com.google.android.googlequicksearchbox":{"notification":"3"}}`

	container := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonStr), &container)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Println(container)
	fmt.Println("")
	for key, subMap := range container {
		//fmt.Println(key, " = ", subMap)

		temp := map[string]interface{}{}
		for k, v := range subMap.(map[string]interface{}) {
			temp[k] = v
		}
		temp["notication"] = 98

		container[key] = temp
	}

	fmt.Println(container)
	fmt.Println("")

	resultByte, _ := json.Marshal(container)
	fmt.Println(string(resultByte))


	var allHourData = make(map[string]string, 24)
	allHourData["0"] = "{\\\"com.baidu.BaiduMap\\\":{\\\"notification\\\":\\\"1\\\"}}"
	allHourData["9"] = "{\\\"com.baidu.BaiduMap\\\":{\\\"notification\\\":\\\"1\\\"},\\\"com.google.android.gms\\\":{\\\"notification\\\":\\\"2\\\"}}"

	allHourDataBytes, _ := json.Marshal(allHourData)
	allHourDataStr := string(allHourDataBytes)
	fmt.Println(allHourDataStr)
}
