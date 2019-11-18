package util

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

var host = []string{
	"http://imoc1:9200/",
}

var Client *elastic.Client

//初始化
func init() {
	var err error
	Client, err = elastic.NewClient(elastic.SetURL(host...))
	if err != nil {
		fmt.Printf("create client failed, err: %v", err)
	}
}
