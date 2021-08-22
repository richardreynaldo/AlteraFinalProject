package transaction_detail

import (
	"errors"
	"finalProject/app/middleware"
	"finalProject/business/transaction_detail"
	"finalProject/business/transaction_header"
	controller "finalProject/controllers"
	"finalProject/controllers/transaction_detail/request"
	"finalProject/controllers/transaction_detail/response"
	headReq "finalProject/controllers/transaction_header/request"
	headResp "finalProject/controllers/transaction_header/response"

	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	transactionDetailUsecase transaction_detail.Usecase
	transactionHeaderUsecase transaction_header.Usecase
}

func NewTransactionDetailController(uc transaction_detail.Usecase, hc transaction_header.Usecase) *TransactionDetailController {
	return &TransactionDetailController{
		transactionDetailUsecase: uc,
		transactionHeaderUsecase: hc,
	}
}

func (ctrl *TransactionDetailController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	userId := middleware.GetUser(c).ID
	req := []request.CreateTransactionDetail{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	headReq := headReq.CreateTransactionHeader{}
	newHeadReq := headReq.ToDomain()
	newHeadReq.UserId = userId
	newHeadReq.Status = "pending"

	for _, j := range req {
		newHeadReq.TotalPrice += j.ToDomain().Price
		newHeadReq.TotalQuantity += j.ToDomain().Quantity
	}

	resp, err := ctrl.transactionHeaderUsecase.Store(ctx, newHeadReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, j := range req {
		j.TransactionId = resp.Id
		_, err := ctrl.transactionDetailUsecase.Store(ctx, j.ToDomain())
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return controller.NewSuccessResponse(c, headResp.FromDomain(resp))
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
