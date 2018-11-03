//server.go

package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/go-redis/redis"
	"gopkg.in/gin-gonic/gin.v1/json"
	"gotests/goPython/gen-go/my/test/demo"
	"os"
)

const (
	NetworkAddr = "0.0.0.0:10086"
)

var Client *redis.Client

type RedisServer struct {
}

func (c *RedisServer) Set(s *demo.HourseSell) (err error) {
	value, err := json.Marshal(s)

	cmd := redis.NewStringCmd("hset", "HourseSell", s.Sname, value)

	Client.Process(cmd)
	//cmd = redis.NewStringCmd("set", "a", "b")
	//
	//Client.Process(cmd)
	//fmt.Println(Client.Set(s.Sname, s.Sid, 1*time.Second))
	return nil
}

func (c *RedisServer) Get(name string) (string, error) {
	cmd := redis.NewStringCmd("hget", "HourseSell", name)
	Client.Process(cmd)
	return cmd.Result()
}

func (c *RedisServer) List(callTime int64) (r []*demo.HourseSell, err error) {
	return r, nil
}

func (c *RedisServer) IsNameExist(name string) (r bool, err error) {

	return false, nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RedisServer{}
	processor := demo.NewClassMemberProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
}
