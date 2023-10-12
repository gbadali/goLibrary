package api

import (
	db "goLibrary/db/sqlc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(db *db.Store) *Server {
	server := &Server{store: db}
	router := gin.Default()

	router.Use(cors.Default())

	// Author Routes
	router.POST("/authors", server.createAuthor)
	router.GET("/authors/:id", server.getAuthor)
	router.GET("/authors", server.searchAuthor)

	// Genre Routes
	router.POST("/genres", server.createGenre)
	router.GET("/genres/:id", server.getGenre)
	router.GET("/genres/", server.searchGenre)
	router.GET("/genres", server.listGenres)

	// Publisher Routes
	router.POST("/publishers", server.createPublisher)
	router.GET("/publishers/:id", server.getPublisher)
	router.GET("/publishers/", server.searchPublisher)
	router.GET("/publishers", server.listPublishers)

	// BookTx Routes
	router.POST("/books", server.createBook)
	router.GET("/books/:id", server.getBook)
	router.GET("/books/", server.listBooks)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
