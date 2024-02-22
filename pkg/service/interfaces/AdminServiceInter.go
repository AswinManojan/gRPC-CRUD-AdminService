package interfaces

import (
	pb "github.com/grpc/gobus/admin/pkg/PB"
	"github.com/grpc/gobus/admin/pkg/model"
)

type AdminServiceInter interface {
	AdminLogin(*pb.LoginRequest) (*model.Admin, error)
	EditUser(*pb.UserDataRequest) (*pb.Result,error)
	DeleteUserById(*pb.IdRequest) (*pb.Result,error)
	FindUserById(*pb.IdRequest) (*pb.Result,error)
	CreateUser(*pb.UserDataRequest) (*pb.Result,error);
	EditProvider(*pb.ProviderDataRequest) (*pb.Result,error)
	FindProviderById(*pb.IdRequest) (*pb.Result,error)
	DeleteProviderById(*pb.IdRequest) (*pb.Result,error)
	CreateProvider(*pb.ProviderDataRequest) (*pb.Result,error);
}
