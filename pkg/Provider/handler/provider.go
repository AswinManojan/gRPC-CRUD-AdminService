package handler

import (
	"context"
	"log"

	adminpb "github.com/grpc/gobus/admin/pkg/PB"
	providerpb "github.com/grpc/gobus/admin/pkg/Provider/pb"
)

func EditProvider(client providerpb.ProviderServicesClient, userData *adminpb.ProviderDataRequest) (*providerpb.Res, error) {
	ctx := context.Background()
	response, err := client.EditProvider(ctx, &providerpb.ProviderDataReq{
		UserName: userData.UserName,
		Email:    userData.Email,
		Password: userData.Password,
	})
	if err != nil {
		log.Printf("error while editing provider")
		return nil, err
	}
	return response, nil
}

func DeleteProviderById(client providerpb.ProviderServicesClient, userData *adminpb.IdRequest) (*providerpb.Res, error) {
	ctx := context.Background()
	response, err := client.DeleteProviderById(ctx, &providerpb.IdReq{
		Id: userData.Id,
	})
	if err != nil {
		log.Printf("error while finding provider")
		return nil, err
	}
	return response, nil
}
func FindProviderById(client providerpb.ProviderServicesClient, userData *adminpb.IdRequest) (*providerpb.Res, error) {
	ctx := context.Background()
	response, err := client.FindProviderById(ctx, &providerpb.IdReq{
		Id: userData.Id,
	})
	if err != nil {
		log.Printf("error while finding provider")
		return nil, err
	}
	return response, nil
}
func CreateProvider(client providerpb.ProviderServicesClient, userData *adminpb.ProviderDataRequest) (*providerpb.Res, error) {
	ctx := context.Background()
	response, err := client.CreateProvider(ctx, &providerpb.ProviderDataReq{
		UserName: userData.UserName,
		Email:    userData.Email,
		Password: userData.Password,
	})
	if err != nil {
		log.Printf("error while creating provider")
		return nil, err
	}
	return response, nil
}
