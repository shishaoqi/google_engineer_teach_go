package main

import (
"crypto/md5"
"fmt"
"io"
)

const (
	TOKEN_SALT1 = "iCX0C0bj59xt"
	TOKEN_SALT2 = "g9C1Hr6O1WfR"
)

/*
* Generate Token for device.
* Note:
*      Given same $device_id, token must be same.
*      Every token is generated by fixed algorithm.
* @param device_id
* @return token
 */
func GenerateToken(device_id string) string {
	return getMD5(getMD5(device_id+TOKEN_SALT1) + TOKEN_SALT2)
}

func getMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	temp := h.Sum(nil)
	return fmt.Sprintf("%x", temp)
}
