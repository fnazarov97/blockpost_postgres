package handlers

import (
	"article/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v2/author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := h.IM.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := h.IM.GetAuthorByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | Created",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/author/{id} [get]
func (h Handler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	author, err := h.IM.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "search"
// @Success     200    {object} models.JSONResponse{data=[]models.Author}
// @Router      /v2/author [get]
func (h Handler) GetAuthorList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Conf.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Conf.DefaultLimit)
	search := c.DefaultQuery("search", "")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "offset error",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "limit error",
		})
		return
	}
	authorList, err := h.IM.GetAuthorList(offset, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    authorList,
	})
}

// UpdateAuthor godoc
// @Summary     update author
// @Description update a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModel true "author body"
// @Success     200    {object} models.JSONResponse{data=models.Author}
// @Response    400    {object} models.JSONErrorResponse
// @Router      /v2/author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var author models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	err := h.IM.UpdateAuthor(author)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Faild update!"})
		return
	}
	updated, err := h.IM.GetAuthorByID(author.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Author not found after updated!"})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "author | Update",
	})
}

// DeleteAuthor godoc
// @Summary     delete author by id
// @Description delete an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "author ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v2/author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	author, err := h.IM.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "This author's id not found!"})
		return
	}
	err = h.IM.DeleteAuthor(author.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Author have been deleted already!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "author deleted",
		"data":    author,
	})
}
