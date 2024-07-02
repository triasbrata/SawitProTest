package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/constants"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

// GetEstateIdDronePlan implements generated.ServerInterface.
func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id string, params generated.GetEstateIdDronePlanParams) error {
	//get data from db
	c := ctx.Request().Context()
	estate, err := s.Repository.GetEstate(c, repository.GetEstateRequest{
		ID: id,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	if len(estate.Data) != 1 {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{StatusCode: int64(http.StatusNotFound), Success: false, Message: fmt.Sprintf("estate with id %s is not found", id)})

	}

	trees, err := s.Repository.GetEstateTree(c, repository.GetEstateTreeRequest{
		EstateID: id,
	})

	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	maxDistance := -1
	if params.MaxDistance != nil {
		maxDistance = *params.MaxDistance
	}
	calc := calculateDronePlan(estate.Data[0], trees.Data, maxDistance)
	return ctx.JSON(http.StatusFound, calc)
}

// GetEstateIdStats implements generated.ServerInterface.
func (s *Server) GetEstateIdStats(ctx echo.Context, id string) error {
	//get data from database
	c := ctx.Request().Context()
	estate, err := s.Repository.GetEstate(c, repository.GetEstateRequest{
		ID: id,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	if len(estate.Data) != 1 {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{StatusCode: int64(http.StatusNotFound), Success: false, Message: fmt.Sprintf("estate with id %s is not found", id)})

	}
	res, err := s.Repository.GetEstateTree(c, repository.GetEstateTreeRequest{
		EstateID: id,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}

	return ctx.JSON(http.StatusAccepted, calculateStats(res))
}

// PostEstate implements generated.ServerInterface.
func (s *Server) PostEstate(ctx echo.Context) error {

	input := model.EstateRequest{}

	//bind request to input and validate it
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: err.Error()})
	}
	if input.Width < constants.MIN_ESTATE_WIDTH_LENGTH || input.Width > constants.MAX_ESTATE_WIDTH_LENGTH {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: fmt.Sprintf("%s of estate max is only %v and min is %v", "Width", constants.MAX_ESTATE_WIDTH_LENGTH, constants.MIN_ESTATE_WIDTH_LENGTH)})
	}
	if input.Length < constants.MIN_ESTATE_WIDTH_LENGTH || input.Length > constants.MAX_ESTATE_WIDTH_LENGTH {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: fmt.Sprintf("%s of estate max is only %v and min is %v", "Leght", constants.MAX_ESTATE_WIDTH_LENGTH, constants.MIN_ESTATE_WIDTH_LENGTH)})
	}
	//transaction to database
	c := ctx.Request().Context()
	res, err := s.Repository.DoCreateEstate(c, repository.DoCreateEstateRequest{
		Length: int64(input.Length),
		Width:  int64(input.Width),
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, model.EstateResponseSuccess{ID: res.ID})
}

// PostEstateIdTree implements generated.ServerInterface.
func (s *Server) PostEstateIdTree(ctx echo.Context, id string) error {

	input := model.TreeRequest{}

	//bind request to input and validate it
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: err.Error()})
	}

	if input.Height < 1 || input.Height > 30 {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: "height of tree max is only 30"})
	}

	if input.X < 1 || input.Y < 1 {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: "out of bound"})
	}

	c := ctx.Request().Context()
	estate, err := s.Repository.GetEstate(c, repository.GetEstateRequest{
		ID: id,
	})

	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	if len(estate.Data) == 0 {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusNotFound), Success: false, Message: "estate not found"})
	}

	if input.X > estate.Data[0].Length && input.Y > estate.Data[0].Width {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{StatusCode: int64(http.StatusBadRequest), Success: false, Message: "out of bound"})
	}
	//transaction to database
	res, err := s.Repository.DoCreateTree(c, repository.DoCreateTreeRequest{
		EstateID: id,
		X:        input.X,
		Y:        input.Y,
		Height:   input.Height,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{StatusCode: int64(http.StatusInternalServerError), Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, model.TreeResponseSuccess{ID: res.ID})

}
