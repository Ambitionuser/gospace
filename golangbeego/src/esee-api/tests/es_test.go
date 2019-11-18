package test

import (
	"encoding/json"
	"esee-api/entry/vo"
	"esee-api/server"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPingNode(t *testing.T) {
	server.PingNode()
}

func TestIndexExists(t *testing.T) {
	result := server.IndexExists("imoc")
	fmt.Println("all index exists: ", result)
}

func TestCreateIndex(t *testing.T) {
	result := server.CreateIndex("esee", vo.Mapping)
	fmt.Println("mapping created: ", result)
}

func TestDeleteIndex(t *testing.T) {
	result := server.DelIndex("esee")
	fmt.Println("all index deleted: ", result)
}

//

//
func TestBatch(t *testing.T) {
	tweet1 := vo.Tweet{User: "Jame1", Age: 23, Message: "Take One", Retweets: 1, Created: time.Now()}
	tweet2 := vo.Tweet{User: "Jame2", Age: 32, Message: "Take Two", Retweets: 0, Created: time.Now()}
	tweet3 := vo.Tweet{User: "Jame3", Age: 32, Message: "Take Three", Retweets: 0, Created: time.Now()}
	server.Batch("esee", "ci", tweet1, tweet2, tweet3)
}

func TestGetDoc(t *testing.T) {
	var tweet vo.Tweet
	data := server.GetDoc("esee", "1")
	if err := json.Unmarshal(data, &tweet); err == nil {
		fmt.Printf("data: %v\n", tweet)
	}
}

//通过字段值查询
func TestTermQuery(t *testing.T) {
	var tweet vo.Tweet
	result := server.TermQuery("esee", "ci", "user", "jame2")
	//获得数据, 方法一
	for _, item := range result.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(vo.Tweet); ok {
			fmt.Printf("tweet : %v\n", t)
		}
	}
}

//
func TestSearch(t *testing.T) {
	result := server.SearchQuery("esee", "ci")
	fmt.Println(result)
	var tweet vo.Tweet
	for _, item := range result.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(vo.Tweet); ok {
			fmt.Printf("tweet : %v\n", t)
		}
	}
}

//
func TestAggsSearch(t *testing.T) {
	server.AggsSearch("esee", "ci")
}
