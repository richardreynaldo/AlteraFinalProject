package articles

import (
	"errors"
	"finalProject/app/middleware"
	"finalProject/business/articles"
	controller "finalProject/controllers"
	"finalProject/controllers/articles/request"
	"finalProject/controllers/articles/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUsecase articles.Usecase
}

func NewArticleController(uc articles.Usecase) *ArticleController {
	return &ArticleController{
		articleUsecase: uc,
	}
}

func (ctrl *ArticleController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	userId := middleware.GetUser(c).ID
	req := request.CreateArticle{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.articleUsecase.Store(ctx, req.ToDomain(), userId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ArticleController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	userId := middleware.GetUser(c).ID
	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.CreateArticle{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainReq := req.ToDomain()
	domainReq.UserId = userId
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt
	resp, err := ctrl.articleUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *ArticleController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.articleUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Article{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
