package db

import (
	"context"
	"database/sql"
	"fmt"
)

type CreateTxBookParams struct {
	Title           string         `json:"title"`
	Isbn            string         `json:"isbn"`
	PublicationYear sql.NullInt32  `json:"publication_year"`
	Author          Author         `json:"author"`
	Publisher       Publisher      `json:"publisher"`
	Genre           Genre          `json:"genre"`
	Price           sql.NullString `json:"price"`
	StockQuantity   sql.NullInt32  `json:"stock_quantity"`
}

type CreateTxBookResults struct {
	Book      Book      `json:"book"`
	Author    Author    `json:"author"`
	Publisher Publisher `json:"publisher"`
	Genre     Genre     `json:"genre"`
}

// CreateTxBook creates a book object but also checks if the Author, Genre and Publisher Exhist and
// creat them if they don't
func (store *Store) CreateTxBook(ctx context.Context, arg CreateTxBookParams) (CreateTxBookResults, error) {
	var result CreateTxBookResults
	fmt.Println(arg)
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// Check if Author exists and add it if it doesn't
		result.Author, err = q.SearchAuthors(ctx, SearchAuthorsParams{
			FirstName: arg.Author.FirstName,
			LastName:  arg.Author.LastName,
		})
		if err == sql.ErrNoRows {
			err = nil
			result.Author, err = q.CreateAuthor(ctx, CreateAuthorParams{
				FirstName: arg.Author.FirstName,
				LastName:  arg.Author.LastName,
			})
			if err != nil {
				return err
			}
		} else {
			result.Author.FirstName = arg.Author.FirstName
			result.Author.LastName = arg.Author.LastName
		}
		if err != nil {
			return err
		}
		// Check if Publisher exists and add if it doesn't
		result.Publisher.PublisherID, err = q.SearchPublisher(ctx, arg.Publisher.PublisherName)
		if err == sql.ErrNoRows {
			err = nil
			result.Publisher, err = q.CreatePublisher(ctx, CreatePublisherParams{
				PublisherName: arg.Publisher.PublisherName,
				Address:       arg.Publisher.Address,
				Phone:         arg.Publisher.Phone,
			})
			if err != nil {
				return err
			}
		} else {
			result.Publisher.PublisherName = arg.Publisher.PublisherName
			result.Publisher.Address = arg.Publisher.Address
			result.Publisher.Phone = arg.Publisher.Phone
		}
		if err != nil {
			return err
		}
		// Check if Genre exists and add it if it doesn't
		result.Genre.GenreID, err = q.SearchGenres(ctx, arg.Genre.GenreName)
		if err == sql.ErrNoRows {
			err = nil
			result.Genre, err = q.CreateGenre(ctx, arg.Genre.GenreName)
			if err != nil {
				return err
			}
		} else {
			result.Genre.GenreName = arg.Genre.GenreName
		}
		if err != nil {
			return err
		}
		// Add the book
		result.Book, err = q.CreateBook(ctx, CreateBookParams{
			Title:           arg.Title,
			Isbn:            arg.Isbn,
			PublicationYear: arg.PublicationYear,
			AuthorID:        sql.NullInt32{Int32: int32(result.Author.AuthorID), Valid: true},
			PublisherID:     sql.NullInt32{Int32: int32(result.Publisher.PublisherID), Valid: true},
			GenreID:         sql.NullInt32{Int32: int32(result.Genre.GenreID), Valid: true},
			Price:           arg.Price,
			StockQuantity:   arg.StockQuantity,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
