package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/json-iterator/go"
	"strconv"
	"time"
	zhugeRedis "zhuge-go/pkg/models/redis"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 入参结构体
type ExecFunc struct {
	Fn      func(map[string]interface{}) (interface{}, error)
	Key     string
	Field   string
	Param   map[string]interface{}
	timeOut int64
}

// 存入redis的结构体
type RedisData struct {
	time int64       `json:"time"`
	data interface{} `json:"data"`
}

// 需要执行的类名和方法 还有项目名称
func Pcache(exeFunc ExecFunc, param map[string]interface{}) interface{} {

	// 拼接大key
	//1.项目名
	//2.类名
	//3.方法名
	key := exeFunc.Key
	// 拼 field
	// 1.cache_key
	field := exeFunc.Field
	//失效时间
	timeOut := exeFunc.timeOut

	// 缓存模式 从pika中获取 默认0  0超时再走缓存，1直接走缓存
	cmd := redis.NewStringCmd("get", "Cache_type")
	zhugeRedis.Redis.RedisClient().Process(cmd)
	d, err := cmd.Result()
	var cacheType int64 = 0

	if err != nil {
		fmt.Print(err)
	} else {
		cacheType, err = strconv.ParseInt(d, 10, 64)
		if err != nil {
			fmt.Print(err)
			cacheType = 0
		}
	}
	// loop步骤
	x := cacheType
	// 返回结果
	var resultData interface{}
	// 执行func  0
	// 正常插入缓存 1
	// 异常读取缓存 2
	// 返回数据 3

	// 读取缓存 2
	// 没有缓存执行func 0
	// 执行func后插入缓存 1
	// 返回数据 3

	//LOOP:for true {
	//var result interface{}
LOOP:
	switch x {
	case 0:
		// 直接执行方法
		resultData, err = exeFunc.Fn(param)
		fmt.Print(err)

		// 有异常 打印异常
		if err != nil {
			fmt.Print(err)
		}

		// 不为空 入缓存 结束
		if resultData != nil {
			x = 2
			goto LOOP
		} else {
			// 没有数据 且缓存类型为0 去读缓存
			if cacheType == 0 {
				x = 1
				goto LOOP
			}
		}
	case 1:
		// 读缓存
		resultData = getCache(key, field, timeOut)
		if cacheType == 1 {
			x = 0
			goto LOOP
		}

	case 2:
		// 插入缓存
		setCache(key, field, resultData)
	default:
		// 不应该走到这啊 直接结束吧 ：）

	}
	return resultData
}

// 读取缓存
func getCache(key string, field string, timeOut int64) interface{} {
	// 获取redis连接对象
	// 读取缓存
	// 是否判断失效
	cmd := redis.New

	StringCmd("hget", field, key)
	zhugeRedis.Redis.RedisClient().Process(cmd)
	redisStringData, err := cmd.Result()
	if err != nil {
		fmt.Print(err)
		redisStringData = ""
	}
	var redisData RedisData
	err1 := json.UnmarshalFromString(redisStringData, &redisData)
	if err1 != nil {
		fmt.Print(err1)
		return nil
	}
	var currentTime = time.Now().Unix()
	if redisData.time+timeOut >= currentTime {
		return redisData.data
	}
	return nil
}

// 读取缓存
func setCache(key string, field string, resultData interface{}) (string, error) {

	// 拼装结构体
	redisData := &RedisData{time.Now().Unix(), resultData}
	// 转为 string
	redisStrData, err := json.Marshal(redisData)

	// 异常直接返回
	if err != nil {
		fmt.Print(err)
		return "0", err
	}

	// 放入redis并返回结果
	cmd := redis.NewStringCmd("hset", field, key, redisStrData)
	zhugeRedis.Redis.RedisClient().Process(cmd)
	return cmd.Result()
}
