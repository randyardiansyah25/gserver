package main

import (
	"context"
	"fmt"
	"gserver/common"
	"log"
	"net"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUserHandlerServer() common.UserHandlerServer {
	return &Server{}
}

type Server struct {
}

func (s *Server) AddUser(context.Context, *common.User) (resp *common.Response, er error) {
	return
}

func (s *Server) GetUser(context.Context, *common.UserId) (resp *common.User, er error) {
	return
}

func (s *Server) ListUser(context.Context, *emptypb.Empty) (resp *common.UserList, er error) {
	list := make([]*common.User, 0)
	r := &common.UserList{}
	fmt.Println("new call : ListUser method...")
	for i := 1; i <= 10; i++ {
		user := common.User{
			Id:       gofakeit.DigitN(8),
			Name:     gofakeit.Name(),
			Password: gofakeit.Password(true, true, true, false, false, 6),
			Gender:   common.UserGender_Male,
		}
		list = append(list, &user)
	}
	r.List = list
	return r, er
}

func main() {
	listener, er := net.Listen("tcp", ":2200")
	if er != nil {
		log.Fatal(er.Error())
	}

	var opts []grpc.ServerOption
	gsvr := grpc.NewServer(opts...)
	common.RegisterUserHandlerServer(gsvr, NewUserHandlerServer())
	fmt.Println("Server listen at 2200")
	gsvr.Serve(listener)
}
