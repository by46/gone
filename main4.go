package main

import (
	"fmt"

	"github.com/by46/gone/im"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/gommon/log"
)

func main() {
	info := &im.HelloWorld{
		Str: proto.String("hello world"),
		Id:  proto.Int32(12),
	}

	data, err := proto.Marshal(info)
	if err != nil {
		log.Fatalf("marshal message error %v", err)
	}
	info2 := &im.HelloWorld{}

	err = proto.Unmarshal(data, info2)
	if err != nil {
		log.Fatalf("Unmarshal message error  %v", err)
	}
	fmt.Printf("hell world id %d, %s", info2.GetId(), info2.GetStr())

	info3 := &im.Author{}
	info3.Ages = append(info3.Ages, 1)
	info4 := &im.Info{}
	info4.TestOneof = &im.Info_Name{Name: "hello"}
	info5 := info4.GetTestOneof()
	switch info5.(type) {
	case *im.Info_Name:
		fmt.Printf("Name fields setting %v", info4.GetName())
	default:
		print("default")
	}

	info6 := &im.DFIS{}
	x, xx := info6.Descriptor()
}
