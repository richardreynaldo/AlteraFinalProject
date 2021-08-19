package reviews

import (
	"errors"
	"finalProject/app/middleware"
	"finalProject/business/reviews"
	controller "finalProject/controllers"
	"finalProject/controllers/reviews/request"
	"finalProject/controllers/reviews/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ReviewController struct {
	reviewUsecase reviews.Usecase
}

func NewReviewController(uc reviews.Usecase) *ReviewController {
	return &ReviewController{
		reviewUsecase: uc,
	}
}

func (ctrl *ReviewController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	userId := middleware.GetUser(c).ID
	req := request.CreateReview{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.reviewUsecase.Store(ctx, req.ToDomain(), userId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ReviewController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.CreateReview{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt
	resp, err := ctrl.reviewUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *ReviewController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.reviewUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Review{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *ReviewController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ctrl.reviewUsecase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	responseController := response.FromDomain(res)
	return controller.NewSuccessResponse(c, responseController)
}
