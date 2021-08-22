package transaction_header

import (
	"errors"
	"finalProject/business/transaction_header"
	controller "finalProject/controllers"
	"finalProject/controllers/transaction_header/request"
	"finalProject/controllers/transaction_header/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type TransactionHeaderController struct {
	transactionHeaderUsecase transaction_header.Usecase
}

func NewTransactionHeaderController(uc transaction_header.Usecase) *TransactionHeaderController {
	return &TransactionHeaderController{
		transactionHeaderUsecase: uc,
	}
}

func (ctrl *TransactionHeaderController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.CreateTransactionHeader{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.transactionHeaderUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TransactionHeaderController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.CreateTransactionHeader{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt
	resp, err := ctrl.transactionHeaderUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *TransactionHeaderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.transactionHeaderUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.TransactionHeader{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
