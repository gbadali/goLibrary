package api

import (
	"database/sql"
	db "goLibrary/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePublisherRequest struct {
	PublisherName string `json:"publisher_name" binding:"required"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
}

// createPublisher is an api endpoint to make an publisher.
func (server *Server) createPublisher(ctx *gin.Context) {
	var req CreatePublisherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	publisher, err := server.store.CreatePublisher(ctx, db.CreatePublisherParams{
		PublisherName: req.PublisherName,
		Address:       sql.NullString{String: req.Address, Valid: false},
		Phone:         sql.NullString{String: req.Phone, Valid: false},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, publisher)
}

type GetPublisherRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getPublisher gets an publisher by id
func (server *Server) getPublisher(ctx *gin.Context) {
	var req GetPublisherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	publisher, err := server.store.GetPublisher(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, publisher)

}

type SearchPublisherRequest struct {
	PublisherName string `json:"publisher_name"`
}

// searchPublisher searches an publisher by firstName lastName
func (server *Server) searchPublisher(ctx *gin.Context) {
	var req SearchPublisherRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	publisher, err := server.store.SearchPublisher(ctx, req.PublisherName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, publisher)
}

// listPublishers lists all the publishers
func (server *Server) listPublishers(ctx *gin.Context) {
	publishers, err := server.store.ListPublishers(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusFound, publishers)
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, publishers)
}
