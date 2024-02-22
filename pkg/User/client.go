package user

import (
	"log"

	"github.com/grpc/gobus/admin/config"
	"github.com/grpc/gobus/admin/pkg/User/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClientDial(cfg *config.Config) (pb.UserServicesClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCUSERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCUSERPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Client at port: %v", cfg.GRPCUSERPORT)
	return pb.NewUserServicesClient(grpc), nil
}
