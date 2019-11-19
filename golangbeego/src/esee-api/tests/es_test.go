package test

import (
	"context"
	"encoding/json"
	"esee-api/entry/vo"
	"esee-api/server"
	"esee-api/util"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
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
	tweet2 := vo.Tweet{User: "Jame2", Age: 1, Message: "Take Two", Retweets: 0, Created: time.Now()}
	tweet3 := vo.Tweet{User: "Jame3", Age: 2, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet4 := vo.Tweet{User: "Jame3", Age: 3, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet5 := vo.Tweet{User: "Jame3", Age: 4, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet6 := vo.Tweet{User: "Jame3", Age: 5, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet7 := vo.Tweet{User: "Jame3", Age: 6, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet8 := vo.Tweet{User: "Jame3", Age: 7, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet9 := vo.Tweet{User: "Jame3", Age: 8, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet10 := vo.Tweet{User: "Jame3", Age: 32, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet11 := vo.Tweet{User: "Jame3", Age: 329, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet12 := vo.Tweet{User: "Jame3", Age: 302, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet13 := vo.Tweet{User: "Jame3", Age: 3027, Message: "Take Three", Retweets: 0, Created: time.Now()}
	tweet14 := vo.Tweet{User: "Jame3", Age: 3072, Message: "Take Three", Retweets: 0, Created: time.Now()}
	server.Batch("esee", "ci", tweet1, tweet2, tweet3, tweet4,
		tweet5, tweet6, tweet7, tweet8, tweet9, tweet10, tweet11, tweet12, tweet13, tweet14)
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

//查找
func TestGets(t *testing.T) {
	//取所有
	var res *elastic.SearchResult
	var err error
	res, err = util.Client.Search("esee").Type("ci").Do(context.Background())
	printEmployee(res, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ vo.Tweet
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(vo.Tweet)
		fmt.Printf("%#v\n", t)
	}
}
