package handler

import (
	"context"
	"fmt"
	"log"

	adminpb "github.com/grpc/gobus/admin/pkg/PB"
	userpb "github.com/grpc/gobus/admin/pkg/User/pb"
)

func EditUser(client userpb.UserServicesClient, userData *adminpb.UserDataRequest) (*userpb.RESULT, error) {
	ctx := context.Background()
	response, err := client.EditUser(ctx, &userpb.UserDataReq{
		UserName: userData.UserName,
		Email:    userData.Email,
		Password: userData.Password,
	})
	if err != nil {
		log.Printf("error while editing user")
		return nil, err
	}
	return response, nil
}

func DeleteUserById(client userpb.UserServicesClient, userData *adminpb.IdRequest) (*userpb.RESULT, error) {
	ctx := context.Background()
	response, err := client.DeleteUserById(ctx, &userpb.IdREQ{
		Id: userData.Id,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}
func FindUserById(client userpb.UserServicesClient, userData *adminpb.IdRequest) (*userpb.RESULT, error) {
	ctx := context.Background()
	response, err := client.FindUserById(ctx, &userpb.IdREQ{
		Id: userData.Id,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}
func CreateUser(client userpb.UserServicesClient, userData *adminpb.UserDataRequest) (*userpb.RESULT, error) {
	ctx := context.Background()
	fmt.Println("----------------------")
	response, err := client.CreateUser(ctx, &userpb.UserDataReq{
		UserName: userData.UserName,
		Email:    userData.Email,
		Password: userData.Password,
	})
	fmt.Println("never")
	if err != nil {
		log.Printf("error while creating user")
		return nil, err
	}
	return response, nil
}
