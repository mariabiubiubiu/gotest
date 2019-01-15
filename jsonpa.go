package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"strings"
)

var term = []string{"cnt", "cityarea_id", "id", "borough_id"}

var like = []string{"tag", "source_name", "borough_address"}

var ranges = []string{"house_price", "house_totalarea", "house_built_year", "created_time",
	"min_price", "house_room"}

var wildcards = []string{"edit_tag", "full_text", "borough_name", "house_title", "cityarea_name", "cityarea2_name",
	"boroug  h_name"}

type Bools struct {
	Must    []map[string]map[string]interface{} `json:"must"`
	MustNot []map[string]map[string]interface{} `json:"must_not"`
	Should  []map[string]map[string]interface{} `json:"should"`
}

type WarpBool struct {
	Must    []map[string]interface{} `json:"must"`
	MustNot []map[string]interface{} `json:"must_not"`
	Should  []map[string]interface{} `json:"should"`
}


func main() {
	var parseMap map[string]interface{}
	buf := []byte(`{"price_updated":"1-100000000000000000"}`)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(buf, &parseMap)
	query := initialize()
	output, _ := json.Marshal(query)
	fmt.Println(string(output))


}


func keyInSli(destMap map[string]interface{}, slices []string, name string, ret *WarpBool) *WarpBool {
	//遍历预设数组
	for _, value := range slices {
		//判断传输过来的字段是否有预设数组中的字段
		if mapValue, ok := destMap[value]; ok {

			term := make(map[string]interface{})
			tempBool := new(Bools)
			//判断是否为not-前缀，若是则按照非策略执行
			if strings.HasPrefix(value, "not-") {
				array := strings.Split(value, "-")
				//判断字段是否要查询范围
				if name != "range" {
					//判断非范围字段term（tempmap）MustNot（MustNot队列）mapValue（传输来的字段值），array[1]（传输来的字段），name（队列名称）
					tempBool.MustNot = judegNotRange(term, tempBool.MustNot, mapValue, array[1], name)
				} else {
					//判断非范围字段term（tempmap）MustNot（MustNot队列）mapValue（传输来的字段值），array[1]（传输来的字段），name（队列名称）
					tempBool.MustNot = judegRange(term, tempBool.MustNot, mapValue, array[1], name)
				}
				//最外层的返回值
				ret.MustNot = append(ret.MustNot, map[string]interface{}{"bool": tempBool})
			} else {
				if name != "range" {
					tempBool.Must = judegNotRange(term, tempBool.Must, mapValue, value, name)
				} else {
					if strings.Contains(mapValue.(string), ",") {
						for _, shouldValue := range strings.Split(mapValue.(string), ",") {
							rangeTemp := make(map[string]interface{})
							tempBool.Should = judegRange(rangeTemp, tempBool.Should, shouldValue, value, name)
						}
					} else {
						tempBool.Should = judegRange(term, tempBool.Should, mapValue, value, name)
					}
				}
				ret.Must = append(ret.Must, map[string]interface{}{"bool": tempBool})
			}
		}
	}
	return ret
}

func judegRange(term map[string]interface{}, tempBool []map[string]map[string]interface{}, mapValue interface{}, value, name string) []map[string]map[string]interface{} {
	temp := make(map[string]interface{})
	if !strings.Contains(mapValue.(string), "-") {
		term[value] = mapValue
		tempBool = append(tempBool, map[string]map[string]interface{}{"term": term})
		return tempBool
	} else {
		ranges := strings.Split(mapValue.(string), "-")
		temp["gte"] = ranges[0]
		temp["lte"] = ranges[1]
		term[value] = temp
		tempBool = append(tempBool, map[string]map[string]interface{}{name: term})
		return tempBool
	}
}

func judegNotRange(term map[string]interface{}, tempBool []map[string]map[string]interface{}, mapValue interface{}, value, name string) []map[string]map[string]interface{} {
	if strings.Contains(mapValue.(string), ",") {
		valueList := strings.Split(mapValue.(string), ",")
		term[value] = valueList
		tempBool = append(tempBool, map[string]map[string]interface{}{"terms": term})
		return tempBool
	} else {
		if name == "wildcard" {
			term[value] = "*" + mapValue.(string) + "*"
			tempBool = append(tempBool, map[string]map[string]interface{}{"wildcard": term})
		} else if name == "match" {
			term[value] = map[string]interface{}{"query": mapValue, "slop": 10}
			tempBool = append(tempBool, map[string]map[string]interface{}{"match_phrase": term})
		} else {
			term[value] = mapValue
			tempBool = append(tempBool, map[string]map[string]interface{}{"term": term})
		}
		return tempBool
	}
}


func initialize()map[string]interface{}{
	warp := new(WarpBool)
	query := make(map[string]interface{})
	query["query"] = warp
	return query
}