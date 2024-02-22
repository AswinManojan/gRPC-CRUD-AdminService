package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc/gobus/admin/config"
	pb "github.com/grpc/gobus/admin/pkg/PB"
	"github.com/grpc/gobus/admin/pkg/handler"
	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config, svc *handler.AdminHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":"+cfg.GRPCADMINPORT)
	if err != nil {
		log.Fatal("error creating listener on port 8083")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, svc)
	fmt.Println("finally")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")
	}
}
