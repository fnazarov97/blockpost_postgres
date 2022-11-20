package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"article/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "article body"
// @Success     201     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article [post]
func (h Handler) CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.IM.AddArticle(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.IM.GetArticleByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "getarticleErr",
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | Created",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/article/{id} [get]
func (h Handler) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	article, err := h.IM.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "search"
// @Success     200    {object} models.JSONResponse{data=[]models.Article}
// @Router      /v2/article [get]
func (h Handler) GetArticleList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Conf.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Conf.DefaultLimit)
	searchStr := c.DefaultQuery("search", "")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "offset error",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "limit error",
		})
		return
	}
	articleList, err := h.IM.GetArticleList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// UpdateArticle godoc
// @Summary     update article
// @Description update a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.UpdateArticleModel true "article body"
// @Success     200     {object} models.JSONResponse{data=[]models.Article}
// @Response    400     {object} models.JSONErrorResponse
// @Router      /v2/article [put]
func (h Handler) UpdateArticle(c *gin.Context) {
	var article models.UpdateArticleModel
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	err := h.IM.UpdateArticle(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	updated, err := h.IM.GetArticleByID(article.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("not found"))
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "Article | Update",
	})
}

// DeleteArticle godoc
// @Summary     delete article by id
// @Description delete an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.DeleteArticleModel}
// @Failure     404 {object} error
// @Router      /v2/article/{id} [delete]
func (h Handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	article, err := h.IM.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Article | Delete | NOT FOUND",
			"data":    nil,
		})
		return
	}
	err = h.IM.DeleteArticle(article.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Article | Delete | NOT FOUND",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article deleted",
		"data":    article,
	})
}
