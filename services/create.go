package services

import (
	"fmt"
	dto "greet-sm-srv/dto/socmed"
)

func (b *blogService) Create(req dto.CreateRequest) (dto.SocmedResponse, error) {

	blog, err := b.SocmedR.Create(req)
	if err != nil {
		return dto.SocmedResponse{}, err
	}

	response := dto.SocmedResponse{
		ID:        req.ID,
		UserID:    req.UserID,
		Image:     req.Image,
		Title:     blog.Title,
		Content:   blog.Content,
		CreatedBy: blog.CreatedBy,
		//
	}
	fmt.Println(response, "b")

	return response, nil
}
