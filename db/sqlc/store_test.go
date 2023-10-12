package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestCreateTxBook(t *testing.T) {
	store := NewStore(testDB)

	gone := Book{
		Title:           "Gone With the Wind",
		Isbn:            "9780582418059",
		PublicationYear: sql.NullInt32{Int32: 1936, Valid: true},
		AuthorID:        sql.NullInt32{Int32: 0, Valid: false},
		PublisherID:     sql.NullInt32{Int32: 0, Valid: false},
		GenreID:         sql.NullInt32{Int32: 0, Valid: false},
		Price:           sql.NullString{String: "9.99", Valid: true},
		StockQuantity:   sql.NullInt32{Int32: 10, Valid: true},
	}

	margo := Author{
		FirstName: "Margret",
		LastName:  "Mitchell",
	}

	histfic := Genre{
		GenreName: "Historical Fiction",
	}

	macmil := Publisher{
		PublisherName: "Macmillan Inc.",
		Address:       sql.NullString{String: "New York", Valid: true},
		Phone:         sql.NullString{String: "555-5432", Valid: true},
	}

	book := CreateTxBookParams{
		Title:           gone.Title,
		Isbn:            gone.Isbn,
		PublicationYear: gone.PublicationYear,
		Author:          margo,
		Publisher:       macmil,
		Genre:           histfic,
		Price:           gone.Price,
		StockQuantity:   gone.StockQuantity,
	}
	fmt.Println("Adding book: ", book)
	results, err := store.CreateTxBook(context.Background(), book)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(results, err)

}
