package server

import (
	"context"
	"esee-api/util"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"strconv"
	"time"
)

//ping 连接测试
func PingNode() {
	start := time.Now()
	info, code, err := util.Client.Ping("http://imoc1:9200/").Do(context.Background())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(code, info.Version.Number)
}

//校验index是否存在
func IndexExists(index ...string) bool {
	exists, err := util.Client.IndexExists(index...).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return exists
}

//创建index
func CreateIndex(index, mapping string) bool {
	result, err := util.Client.CreateIndex(index).BodyString(mapping).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return result.Acknowledged
}

//删除index
func DelIndex(index string) bool {
	response, err := util.Client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return response.Acknowledged
}

//获取指定id的文档
func GetDoc(index, id string) []byte {
	result, err := util.Client.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if result.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", result.Id, result.Version, result.Index, result.Type)
	}
	source, err := result.Source.MarshalJSON()
	if err != nil {
		fmt.Printf("byte convert string failed, err: %v", err)
	}
	return source
}

//批量插入
func Batch(index string, type_ string, datas ...interface{}) {

	bulkRequest := util.Client.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	fmt.Printf("error: %v, %v\n", response.Errors, iter)
}

//term 查询
func TermQuery(index, type_, fieldName, fieldValue string) *elastic.SearchResult {
	query := elastic.NewTermQuery(fieldName, fieldValue)
	//_ = elastic.NewQueryStringQuery(fieldValue) //关键字查询
	searchResult, err := util.Client.Search().
		Index(index).Type(type_).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("query cost %d millisecond.\n", searchResult.TookInMillis)
	return searchResult
}

func SearchQuery(index, type_ string) *elastic.SearchResult {
	boolQuery := elastic.NewBoolQuery()
	//boolQuery.Must(elastic.NewMatchQuery("user", "Jame10"))
	//boolQuery.Filter(elastic.NewRangeQuery("age").Gt("30"))
	searchResult, err := util.Client.Search(index).
		Type(type_).
		Query(boolQuery).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return searchResult
}

func AggsSearch(index, type_ string) {
	minAgg := elastic.NewMinAggregation().Field("age")
	rangeAgg := elastic.NewRangeAggregation().Field("age").AddRange(0, 30).AddRange(30, 60).Gt(60)
	build := util.Client.Search(index).Type(type_).Pretty(true)
	minResult, err := build.Aggregation("minAgg", minAgg).Do(context.Background())
	rangeResult, err := build.Aggregation("rangeAgg", rangeAgg).Do(context.Background())
	if err != nil {
		panic(err)
	}
	minAggRes, _ := minResult.Aggregations.Min("minAgg")
	fmt.Printf("min: %v\n", *minAggRes.Value)
	rangeAggRes, _ := rangeResult.Aggregations.Range("rangeAgg")
	for _, item := range rangeAggRes.Buckets {
		fmt.Printf("key: %s, value: %v\n", item.Key, item.DocCount)
	}
}
