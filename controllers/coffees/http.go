package coffees

import (
	"errors"
	"finalProject/business/coffees"
	controller "finalProject/controllers"
	"finalProject/controllers/coffees/request"
	"finalProject/controllers/coffees/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type CoffeeController struct {
	coffeesUsecase coffees.Usecase
}

func NewCoffeesController(uc coffees.Usecase) *CoffeeController {
	return &CoffeeController{
		coffeesUsecase: uc,
	}
}

func (ctrl *CoffeeController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.CreateCoffee{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.coffeesUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *CoffeeController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.CreateCoffee{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt
	resp, err := ctrl.coffeesUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *CoffeeController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.coffeesUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Coffee{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *CoffeeController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ctrl.coffeesUsecase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	responseController := response.FromDomain(res)
	return controller.NewSuccessResponse(c, responseController)
}
