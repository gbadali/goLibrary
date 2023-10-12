CREATE TABLE "authors" (
    "author_id" bigserial PRIMARY KEY,
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL
);

CREATE TABLE "publishers" (
    "publisher_id" bigserial PRIMARY KEY,
    "publisher_name" VARCHAR(100) NOT NULL,
    "address" VARCHAR(255),
    "phone" VARCHAR(15)
);

CREATE TABLE "genres" (
    "genre_id" bigserial PRIMARY KEY,
    "genre_name" VARCHAR(50) NOT NULL
);

CREATE TABLE "books" (
    "book_id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "isbn" VARCHAR(13) NOT NULL,
    "publication_year" INT,
    "author_id" INT, -- Foreign key referencing Authors table
    "publisher_id" INT, -- Foreign key referencing Publishers table
    "genre_id" INT, -- Foreign key referencing Genres table
    "price" DECIMAL(10, 2),
    "stock_quantity" INT
);

CREATE TABLE "book_authors" (
    "book_author_id" bigserial PRIMARY KEY,
    "book_id" INT, -- Foreign key referencing Books table
    "author_id" INT -- Foreign key referencing Authors table
);

ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("author_id");
ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("publisher_id");
ALTER TABLE "books" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("genre_id");
ALTER TABLE "book_authors" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("book_id");
ALTER TABLE "book_authors" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("author_id");
CREATE INDEX ON "authors" ("first_name");
CREATE INDEX ON "authors" ("last_name");
CREATE INDEX ON "publishers" ("publisher_name");
CREATE INDEX ON "genres" ("genre_name");
CREATE INDEX ON "books" ("title");