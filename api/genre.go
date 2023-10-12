package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGenreRequest struct {
	GenreName string `json:"genre_name" binding:"required"`
}

// createGenre is an api endpoint to make an genre.
func (server *Server) createGenre(ctx *gin.Context) {
	var req CreateGenreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.CreateGenre(ctx, req.GenreName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, genre)
}

type GetGenreRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getGenre gets an genre by id
func (server *Server) getGenre(ctx *gin.Context) {
	var req GetGenreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.GetGenre(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, genre)

}

type SearchGenreRequest struct {
	GenreName string `json:"genre_name"`
}

// searchGenre searches an genre by GenreName
func (server *Server) searchGenre(ctx *gin.Context) {
	var req SearchGenreRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genre, err := server.store.SearchGenres(ctx, req.GenreName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, genre)
}

// listGenres lists all the gennres
func (server *Server) listGenres(ctx *gin.Context) {
	genres, err := server.store.ListGenres(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusFound, genres)
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, genres)
}
