package benchmark

//
//import (
//	"context"
//	"gopkg.in/olivere/elastic.v5"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	"testing"
//)
//
//var host = "http://101.201.119.240:9200"
//var client *elastic.Client
//
//func init() {
//	var err error
//	client, err = elastic.NewClient(
//		elastic.SetURL(host),
//		elastic.SetSniff(false))
//	if err != nil {
//		panic(err)
//	}
//}
//
//func BenchmarkHttp(b *testing.B){
//	b.ReportAllocs()
//	for i := 0;i < b.N; i++{
//		rawResultfulApi()
//	}
//}
//
//
//func BenchmarkRaw(b *testing.B){
//	b.ReportAllocs()
//	for i := 0;i < b.N; i++{
//		search("offline_bj_house_sell", "spider")
//	}
//}
//
//func rawResultfulApi(){
//	var urls = "http://101.201.119.240:9200/offline_bj_house_sell/_search?pretty"
//	//postParam := url.Values{
//	//    "query": {"{\"term\":{\"id\":1}}"},
//	//  }
//	dsl := `{
// "size":0,
// "aggs": {
//   "aggs_borough": {
//     "terms": {
//       "field": "borough_id"
//
//     }
//   }
// }
//}`
//	client := &http.Client{}
//	req, _ := http.NewRequest("POST", urls, strings.NewReader(dsl))
//
//	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
//
//	resp, _ := client.Do(req)
//
//	defer resp.Body.Close()
//	ioutil.ReadAll(resp.Body)
//	//fmt.Println(string(res))
//}
//
//
//
//func search(index_name string, doc_name string) {
//
//	//q := elastic.RawStringQuery(dsl)
//	aggs := elastic.NewTermsAggregation().Field("borough_id")
//	//res, err := client.Search().Index(index_name).Type(doc_name).Query(q).Do(context.Background())
//	client.Search().Index(index_name).Type(doc_name).Aggregation("borough", aggs).Do(context.Background())
//	//fmt.Println(res.Hits.TotalHits)
//	//fmt.Println(res.Aggregations)
//	//for _, item := range  res.Aggregations   {
//	//    fmt.Println(string(*item))
//	//}
//
//}
