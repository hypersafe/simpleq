package main

import (
	// "fmt"
	"github.com/hypersafe/simpleq/message"
	grpc "google.golang.org/grpc"
	"net"
	"log"
	svc "github.com/hypersafe/simpleq/services"
	"context"
	"fmt"

)

type Msgserver struct{}

func( msgsvc *Msgserver) NewMessage(ctx context.Context, in *svc.Msg) (*svc.Status, error){


	msg:= message.NewMsgFromParam(in.Mfrom, in.Mto, in.Mcontent)
	fmt.Println(msg)

	return &svc.Status{Code: 200}, nil

}



func StartServer(address string){

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	svc.RegisterMessageServiceServer(s, &Msgserver{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func main(){
	
	StartServer("localhost:50051")
	
}