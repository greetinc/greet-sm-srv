package repositories

import (
	dto "greet-sm-srv/dto/socmed"
	"greet-sm-srv/entity"

	util "github.com/greetinc/greet-util/s"
)

func (b *blogRepository) Create(req dto.CreateRequest) (*dto.SocmedResponse, error) {
	blog := entity.Socmed{
		ID:        util.GenerateRandomString(),
		Title:     req.Title,
		Content:   req.Content,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	response := &dto.SocmedResponse{
		ID:      blog.ID,
		Title:   blog.Title,
		Content: blog.Content,
		UserID:  req.UserID,
	}
	if err := b.DB.Create(&blog).Error; err != nil {
		return nil, err
	}

	return response, nil
}
