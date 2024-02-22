package interfaces

import "github.com/grpc/gobus/admin/pkg/model"

type AdminRepoInterface interface {
	FetchAdmin(username string) (*model.Admin,error)
}
