//client.go

package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gotests/goPython/gen-go/my/test/demo"
	"net"
	"os"
	"time"
)

const (
	HOST = "127.0.0.1"
	PORT = "10086"
)

func main() {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := demo.NewClassMemberClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	var i int32
	for i = 0; i < 5; i++ {
		var s demo.HourseSell
		s.Sid = i
		s.Sname = fmt.Sprintf("muweijia%d", i)
		err := client.Set(&s)
		if err != nil {
			fmt.Println("err", err)
		}
		//time.Sleep(time.Second * 3)
		fmt.Println(client.Get(s.Sname))
	}

	sList, err := client.List(currentTimeMillis())
	fmt.Println(err)
	for _, s := range sList {
		fmt.Println(s)
	}

	endTime := currentTimeMillis()
	fmt.Printf("calltime:%d-%d=%dms\n", endTime, startTime, (endTime - startTime))

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
