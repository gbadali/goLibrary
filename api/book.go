package api

import (
	"database/sql"
	"fmt"
	db "goLibrary/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookRequest struct {
	Title           string       `json:"title"`
	Isbn            string       `json:"isbn"`
	PublicationYear int32        `json:"publication_year"`
	Author          db.Author    `json:"author"`
	Publisher       db.Publisher `json:"publisher"`
	Genre           db.Genre     `json:"genre"`
	Price           string       `json:"price"`
	StockQuantity   int32        `json:"stock_quantity"`
}

func (server *Server) createBook(ctx *gin.Context) {
	var req CreateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTxBookParams{
		Title:           req.Title,
		Isbn:            req.Isbn,
		PublicationYear: sql.NullInt32{Int32: req.PublicationYear, Valid: true},
		Author:          req.Author,
		Publisher:       req.Publisher,
		Genre:           req.Genre,
		Price:           sql.NullString{String: req.Price, Valid: true},
		StockQuantity:   sql.NullInt32{Int32: req.StockQuantity, Valid: true},
	}

	book, err := server.store.CreateTxBook(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse((err)))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type GetBookRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getGenre gets an genre by id
func (server *Server) getBook(ctx *gin.Context) {
	var req GetBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, book)

}

type ListBooksRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listBooks(ctx *gin.Context) {
	var req ListBooksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Print(req)

	books, err := server.store.ListBooksJoin(ctx, db.ListBooksJoinParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, books)

}
