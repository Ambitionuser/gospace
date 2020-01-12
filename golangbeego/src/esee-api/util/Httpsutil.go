package util

import "fmt"

func GetXtoken(cookie string) string {
	var eiamstatus string
	fmt.Println(len(cookie))
	if len(cookie) < 50 {
		eiamstatus = "500"
	} else {
		eiamstatus = "200"
	}
	return eiamstatus
}
