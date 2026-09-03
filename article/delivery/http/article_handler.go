package http

import (
	"backend/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase) {
	handler := &ArticleHandler{AUsecase: us}

	e.POST("/article/", handler.Store)
	e.GET("/article/:limit/:offset", handler.Fetch)
	e.GET("/article/:id", handler.GetByID)
	e.PUT("/article/:id", handler.Update)
	e.PATCH("/article/:id", handler.Update)
	e.DELETE("/article/:id", handler.Delete)
}

func (a *ArticleHandler) Fetch(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))

	list, err := a.AUsecase.Fetch(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func (a *ArticleHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	art, err := a.AUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "article not found"})
	}
	return c.JSON(http.StatusOK, art)
}

func (a *ArticleHandler) Store(c echo.Context) error {
	var article domain.Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	if err := c.Validate(&article); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := a.AUsecase.Store(c.Request().Context(), &article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, article)
}

func (a *ArticleHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	existing, err := a.AUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "article not found"})
	}

	var article domain.Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	if err := c.Validate(&article); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	article.ID = existing.ID
	err = a.AUsecase.Update(c.Request().Context(), &article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, article)
}

func (a *ArticleHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, err := a.AUsecase.GetByID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "article not found"})
	}

	err := a.AUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully"})
}
