package service

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/grpc/gobus/admin/pkg/PB"
	providerHandler "github.com/grpc/gobus/admin/pkg/Provider/handler"
	providerpb "github.com/grpc/gobus/admin/pkg/Provider/pb"
	userHandler "github.com/grpc/gobus/admin/pkg/User/handler"
	userpb "github.com/grpc/gobus/admin/pkg/User/pb"
	"github.com/grpc/gobus/admin/pkg/model"
	repointerfaces "github.com/grpc/gobus/admin/pkg/repo/Interfaces"
	"github.com/grpc/gobus/admin/pkg/service/interfaces"
)

type AdminService struct {
	repo           repointerfaces.AdminRepoInterface
	userclient     userpb.UserServicesClient
	providerclient providerpb.ProviderServicesClient
}

// CreateProvider implements interfaces.AdminServiceInter.
func (as *AdminService) CreateProvider(providerData *pb.ProviderDataRequest) (*pb.Result, error) {
	result, err := providerHandler.CreateProvider(as.providerclient, providerData)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// CreateUser implements interfaces.AdminServiceInter.
func (as *AdminService) CreateUser(userData *pb.UserDataRequest) (*pb.Result, error) {
	fmt.Println("Hiiiii")
	result, err := userHandler.CreateUser(as.userclient, userData)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// DeleteProviderById implements interfaces.AdminServiceInter.
func (as *AdminService) DeleteProviderById(id *pb.IdRequest) (*pb.Result, error) {
	result, err := providerHandler.DeleteProviderById(as.providerclient, id)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// DeleteUserById implements interfaces.AdminServiceInter.
func (as *AdminService) DeleteUserById(id *pb.IdRequest) (*pb.Result, error) {
	result, err := userHandler.DeleteUserById(as.userclient, id)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// EditProvider implements interfaces.AdminServiceInter.
func (as *AdminService) EditProvider(providerData *pb.ProviderDataRequest) (*pb.Result, error) {
	result, err := providerHandler.EditProvider(as.providerclient, providerData)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// EditUser implements interfaces.AdminServiceInter.
func (as *AdminService) EditUser(userData *pb.UserDataRequest) (*pb.Result, error) {
	result, err := userHandler.EditUser(as.userclient, userData)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// FindProviderById implements interfaces.AdminServiceInter.
func (as *AdminService) FindProviderById(id *pb.IdRequest) (*pb.Result, error) {
	result, err := providerHandler.FindProviderById(as.providerclient, id)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// FindUserById implements interfaces.AdminServiceInter.
func (as *AdminService) FindUserById(id *pb.IdRequest) (*pb.Result, error) {
	result, err := userHandler.FindUserById(as.userclient, id)
	if err != nil {
		return nil, err
	}
	response := &pb.Result{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

// AdminLogin implements interfaces.AdminServiceInter.
func (as *AdminService) AdminLogin(req *pb.LoginRequest) (*model.Admin, error) {
	admin, err := as.repo.FetchAdmin(req.Username)
	if err != nil {
		log.Print("Unable to fetch admin- admin service.")
		return nil, err
	}
	if admin.Password != req.Password {
		log.Print("Password incorrect")
		return nil, errors.New("incorrect Password provided")
	}
	return admin, nil
}

func NewAdminService(REPO repointerfaces.AdminRepoInterface, userClient userpb.UserServicesClient, providerClient providerpb.ProviderServicesClient) interfaces.AdminServiceInter {
	return &AdminService{
		repo:           REPO,
		userclient:     userClient,
		providerclient: providerClient,
	}
}
