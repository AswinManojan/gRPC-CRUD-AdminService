package di

import (
	"log"

	"github.com/grpc/gobus/admin/config"
	db "github.com/grpc/gobus/admin/pkg/DB"
	provider "github.com/grpc/gobus/admin/pkg/Provider"
	user "github.com/grpc/gobus/admin/pkg/User"
	"github.com/grpc/gobus/admin/pkg/handler"
	"github.com/grpc/gobus/admin/pkg/repo"
	"github.com/grpc/gobus/admin/pkg/server"
	"github.com/grpc/gobus/admin/pkg/service"
)

func Init() {
	cfg := config.LoadConfig()
	DB := db.Db_connect(cfg)
	userclient, err := user.UserClientDial(cfg)
	if err != nil {
		log.Fatal("something went wrong-userclient", err)
	}
	providerclient, err := provider.ProviderClientDial(cfg)
	if err != nil {
		log.Fatal("something went wrong-userclient", err)
	}
	repo := repo.NewAdminRepo(DB)
	service := service.NewAdminService(repo, userclient, providerclient)
	handler := handler.NewAdminHandler(service)
	server.NewGRPCServer(cfg, handler)
}
