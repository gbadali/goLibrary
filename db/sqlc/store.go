package db

import (
	"database/sql"
	"encoding/json"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Custom Marshaler to not emit the sql.nullxxx objects.
func (b Book) MarshalJSON() ([]byte, error) {
	bookJSON := make(map[string]interface{})
	bookJSON["book_id"] = b.BookID
	bookJSON["title"] = b.Title
	bookJSON["author_id"] = NullInt32ToJSON(b.AuthorID)
	bookJSON["isbn"] = b.Isbn
	bookJSON["publication_year"] = NullInt32ToJSON(b.PublicationYear)
	bookJSON["genre_id"] = NullInt32ToJSON(b.GenreID)
	bookJSON["price"] = NullStringToJSON(b.Price)
	bookJSON["stock_quantity"] = NullInt32ToJSON(b.StockQuantity)

	return json.Marshal(bookJSON)
}

func (b ListBooksJoinRow) MarshalJSON() ([]byte, error) {
	bookJSON := make(map[string]interface{})
	bookJSON["book_title"] = b.BookTitle
	bookJSON["isbn"] = b.Isbn
	bookJSON["publication_year"] = NullInt32ToJSON(b.PublicationYear)
	bookJSON["author_first_name"] = b.AuthorFirstName
	bookJSON["author_last_name"] = b.AuthorLastName
	bookJSON["publisher_name"] = b.PublisherName
	bookJSON["genre_name"] = b.GenreName
	bookJSON["price"] = NullStringToJSON(b.Price)
	bookJSON["stock_quantity"] = NullInt32ToJSON(b.StockQuantity)

	return json.Marshal(bookJSON)
}

func NullInt32ToJSON(n sql.NullInt32) interface{} {
	if n.Valid {
		return n.Int32
	}
	return nil
}

func NullStringToJSON(n sql.NullString) interface{} {
	if n.Valid {
		return n.String
	}
	return nil
}
