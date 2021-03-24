package main

import (
	"github.com/antsmallant/testgo/aaa"
	"github.com/golang/protobuf/proto"
	pb "github.com/antsmallant/testgo/pbs"
	"fmt"
)

func test_pb() {
	person := &pb.Person{
		Name:"yz",
		Age:11,
		Emails:[]string{"782365461@qq.com","123456@163.com"},
		Phones:[]*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_WORK,
			},
		},
	}
 
	//marshal:  obj---[]byte
	data,err := proto.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
 
	//unmarshal : []byte---obj
	newPersonObj := &pb.Person{}
	err = proto.Unmarshal(data,newPersonObj)
	if err != nil {
		fmt.Println(err)
	}
 
	fmt.Println(newPersonObj)
}

func main() {
	aaa.Pp()

	test_pb()
}