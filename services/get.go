package services

import (
	"fmt"
	dto "greet-sm-srv/dto/socmed"
)

func (s *blogService) GetAll(paginationDTO dto.GetAllRequest) (dto.PaginationResponse, error) {

	data, err := s.SocmedR.GetAll(paginationDTO)
	if err != nil {
		return dto.PaginationResponse{}, err
	}
	// get current url path

	// search query params
	searchQueryParams := ""

	for _, search := range paginationDTO.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	return data, nil
	// dto.Response{Status: true, Data: data}
}

// 	return data, nil
// }
