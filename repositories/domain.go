package repositories

import (
	dto "greet-sm-srv/dto/socmed"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	Create(req dto.CreateRequest) (*dto.SocmedResponse, error)
	GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error)
}

type blogRepository struct {
	DB *gorm.DB
}

func NewSocmedRepository(DB *gorm.DB) SocmedRepository {
	return &blogRepository{
		DB: DB,
	}
}
