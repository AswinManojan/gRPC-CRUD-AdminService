package handler

import (
	"context"
	"fmt"
	"log"

	pb "github.com/grpc/gobus/admin/pkg/PB"
	"github.com/grpc/gobus/admin/pkg/service/interfaces"
)

type AdminHandler struct {
	svc interfaces.AdminServiceInter
	pb.AdminServiceServer
}

func (ah *AdminHandler) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Result, error) {
	_, err := ah.svc.AdminLogin(req)
	if err != nil {
		log.Println("error while logging in")
		return nil, err
	}
	ret := &pb.Result{
		Status:  "Success",
		Error:   "",
		Message: "Admin Login Successful",
	}
	return ret, nil
}

func (a *AdminHandler) FindUserById(ctx context.Context, p *pb.IdRequest) (*pb.Result, error) {
	result, err := a.svc.FindUserById(p)
	if err != nil {
		log.Println("error while finding the user")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) FindProviderById(ctx context.Context, p *pb.IdRequest) (*pb.Result, error) {
	result, err := a.svc.FindProviderById(p)
	if err != nil {
		log.Println("error while finding the provider")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) CreateProvider(ctx context.Context, p *pb.ProviderDataRequest) (*pb.Result, error) {
	result, err := a.svc.CreateProvider(p)
	if err != nil {
		log.Println("error while creating the provider")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) CreateUser(ctx context.Context, p *pb.UserDataRequest) (*pb.Result, error) {
	fmt.Println("Reached")
	result, err := a.svc.CreateUser(p)
	if err != nil {
		log.Println("error while creating the user")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) DeleteUserById(ctx context.Context, p *pb.IdRequest) (*pb.Result, error) {
	result, err := a.svc.DeleteUserById(p)
	if err != nil {
		log.Println("error while deleting the user")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) DeleteProviderById(ctx context.Context, p *pb.IdRequest) (*pb.Result, error) {
	result, err := a.svc.DeleteProviderById(p)
	if err != nil {
		log.Println("error while creating the provider")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) EditUser(ctx context.Context, p *pb.UserDataRequest) (*pb.Result, error) {
	result, err := a.svc.EditUser(p)
	if err != nil {
		log.Println("error while editing the user")
		return nil, err
	}
	return result, nil
}
func (a *AdminHandler) EditProvider(ctx context.Context, p *pb.ProviderDataRequest) (*pb.Result, error) {
	result, err := a.svc.EditProvider(p)
	if err != nil {
		log.Println("error while editing the provider")
		return nil, err
	}
	return result, nil
}
func NewAdminHandler(SVC interfaces.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		svc: SVC,
	}
}
