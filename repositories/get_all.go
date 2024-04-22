package repositories

import (
	"fmt"
	dto "greet-sm-srv/dto/socmed"
	"greet-sm-srv/entity"
	"strings"
)

func (r *blogRepository) GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error) {
	var data []entity.Socmed

	find := r.DB.Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Order(req.SortBy)

	// Menghitung total data yang sesuai dengan pencarian
	var total int64
	if err := r.DB.Model(&entity.Socmed{}).Count(&total).Error; err != nil {
		return dto.PaginationResponse{}, err
	}

	searchs := req.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	if err := find.
		Joins("LEFT JOIN friends ON friends.user_id = socmeds.user_id").
		Find(&data).Error; err != nil {
		return dto.PaginationResponse{}, err
	}

	dataPerPage := req.Limit
	dataPassed := (req.Page - 1) * dataPerPage
	nextPageDataStart := dataPassed + dataPerPage - req.Limit
	nextPage := req.Page + 1
	previousPage := req.Page - 1

	numRows := len(data)

	// Buat respons JSON dengan menggunakan struct PaginationResponse
	response := dto.PaginationResponse{
		TotalData:    int(total),
		TotalRows:    numRows,
		Limit:        req.Limit,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		NextPageData: nextPageDataStart,
		Data:         data,
		Searchs:      searchs,
	}

	return response, nil
}
