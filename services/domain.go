package services

import (
	dto "greet-sm-srv/dto/socmed"
	repositories "greet-sm-srv/repositories"

	m "github.com/greetinc/greet-middlewares/middlewares"
)

type SocmedService interface {
	GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error)
	// GetById(req dto.GetByIdRequest) (*dto.SocmedResponse, error)
	Create(req dto.CreateRequest) (dto.SocmedResponse, error)
	// Update(req dto.UpdateRequest) (dto.SocmedResponse, error)
	// Delete(req dto.DeleteRequest) (dto.SocmedResponse, error)
}

type blogService struct {
	SocmedR repositories.SocmedRepository
	jwt     m.JWTService
}

func NewSocmedService(SocmedR repositories.SocmedRepository, jwtS m.JWTService) SocmedService {
	return &blogService{
		SocmedR: SocmedR,
		jwt:     jwtS,
	}
}
