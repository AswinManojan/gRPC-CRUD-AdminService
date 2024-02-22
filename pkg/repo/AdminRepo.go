package repo

import (
	"log"

	"github.com/grpc/gobus/admin/pkg/model"
	interfaces "github.com/grpc/gobus/admin/pkg/repo/Interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

// FetchAdmin implements interfaces.AdminRepoInterface.
func (ar *AdminRepo) FetchAdmin(username string) (*model.Admin, error) {
	var admin *model.Admin
	result := ar.DB.Where("username=?", username).Find(&admin)
	if result.Error != nil {
		log.Print("Error while fetching admin")
		return nil, result.Error
	}
	return admin, nil
}

func NewAdminRepo(db *gorm.DB) interfaces.AdminRepoInterface {
	return &AdminRepo{
		DB: db,
	}
}
