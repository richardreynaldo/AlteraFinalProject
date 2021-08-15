package transaction_header

import (
	"errors"
	"finalProject/business/transaction_detail"
	controller "finalProject/controllers"
	"finalProject/controllers/transaction_detail/request"
	"finalProject/controllers/transaction_detail/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	transactionDetailUsecase transaction_detail.Usecase
}

func NewArticleController(uc transaction_detail.Usecase) *TransactionDetailController {
	return &TransactionDetailController{
		transactionDetailUsecase: uc,
	}
}

func (ctrl *TransactionDetailController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.CreateTransactionDetail{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.transactionDetailUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TransactionDetailController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.CreateTransactionDetail{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt
	resp, err := ctrl.transactionDetailUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *TransactionDetailController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.transactionDetailUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.TransactionDetail{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
