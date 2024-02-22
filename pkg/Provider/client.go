package provider

import (
	"log"

	"github.com/grpc/gobus/admin/config"
	"github.com/grpc/gobus/admin/pkg/Provider/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProviderClientDial(cfg *config.Config) (pb.ProviderServicesClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCSERVICEPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCSERVICEPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Client at port: %v", cfg.GRPCSERVICEPORT)
	return pb.NewProviderServicesClient(grpc), nil
}
