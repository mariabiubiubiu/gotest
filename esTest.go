package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"time"
)

var host = "http://101.201.119.240:9200"
var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}


func search(index_name string, doc_name string) {

	//q := elastic.RawStringQuery(dsl)
	//aggs := elastic.NewTermsAggregation().Field("borough_id")
	//res, err := client.Search().Index(index_name).Type(doc_name).Query(q).Do(context.Background())
	//client.Search().Index(index_name).Type(doc_name).Aggregation("borough", aggs).Do(context.Background())
	q := elastic.RawStringQuery(`{"bool": {
      "must": null,
      "must_not": null,
      "should": [
        {
          "term": {
            "house_room": "2"
          }
        }
      ]
    }}`)
	result, err:= client.Search().Index(index_name).Type(doc_name).Query(q).From(0).Size(1).Do(context.Background())
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result.Hits.Hits)

	//fmt.Println(res.Hits.TotalHits)
	//fmt.Println(res.Aggregations)
	for _, item := range  result.Hits.Hits   {
	   fmt.Println(item)
	}

}

func main() {
	a := time.Now().UnixNano()
	search("offline_bj_house_sell", "spider")
	b := time.Now().UnixNano()
	fmt.Println("raw", b - a)
}