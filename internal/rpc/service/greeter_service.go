package service

import (
	"context"
	"log"

	"github.com/daheige/goapp/config"
	"github.com/daheige/goapp/internal/dao"
	"github.com/daheige/goapp/pb"
)

// GreeterService greeter service
type GreeterService struct {
	pb.UnimplementedGreeterServiceServer
}

// SayHello say something
// http://localhost:1336/v1/say/1
func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloReply, error) {
	log.Println(config.AppServerConf.HttpPort)
	log.Println("req data: ", in)

	if in.Id > 0 {
		user, err := dao.NewUserDao().GetUser(in.Id)
		if err != nil {
			return nil, err
		}

		return &pb.HelloReply{
			Name:    "username: " + user.Name,
			Message: "call ok",
		}, nil
	}

	return &pb.HelloReply{
		Name:    "hello world",
		Message: "call ok",
	}, nil
}

// Info method.
func (s *GreeterService) Info(ctx context.Context, in *pb.InfoReq) (*pb.InfoReply, error) {
	log.Println(in.Name)

	return &pb.InfoReply{
		Address: "shenzhen",
		Message: "ok",
	}, nil
}
