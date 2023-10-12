package api

import (
	"database/sql"
	db "goLibrary/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAuthorRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// createAuthor is an api endpoint to make an author.
func (server *Server) createAuthor(ctx *gin.Context) {
	var req CreateAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAuthorParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	author, err := server.store.CreateAuthor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type GetAuthorRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAuthor gets an author by id
func (server *Server) getAuthor(ctx *gin.Context) {
	var req GetAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.GetAuthor(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)

}

type SearchAuthorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// searchAuthor searches an author by firstName lastName
func (server *Server) searchAuthor(ctx *gin.Context) {
	var req SearchAuthorRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println(req.FirstName)
	log.Println(req.LastName)

	author, err := server.store.SearchAuthors(ctx, db.SearchAuthorsParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)
}
