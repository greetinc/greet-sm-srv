package handlers

import (
	dto "greet-sm-srv/dto/socmed"

	res "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

func (b *domainHandler) Create(c echo.Context) error {
	var req dto.CreateRequest
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId
	req.CreatedBy = createdBy

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	data, err := b.domainService.Create(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)

}
