package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	GetProducts() ([]*Product, error)
	CreateProduct(*Product) (int64, error)
	UpdateProduct(*Product, int64) error
	DeleteProduct(int64) error
}

type SqliteStore struct {
	db *sql.DB
}

func (s *SqliteStore) CreateProduct(p *Product) (int64, error) {
	query := `insert into product
	(name, price, currency, in_stock)
	values
	($1,$2,$3,$4)`

	resp, err := s.db.Exec(
		query,
		p.Name,
		p.Price,
		p.Currency,
		p.InStock,
	)
	if err != nil {
		return 0, err
	}

	return resp.LastInsertId()
}

func (s *SqliteStore) UpdateProduct(p *Product, id int64) error {
	query := `update product
	set 
		name = $1,
		price = $2,
		currency = $3,
		in_stock = $4
	where
		id = $5
	`

	_, err := s.db.Exec(
		query,
		p.Name,
		p.Price,
		p.Currency,
		p.InStock,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteStore) GetProducts() ([]*Product, error) {
	rows, err := s.db.Query("select * from product where is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	for rows.Next() {
		p, err := scanIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	return products, nil
}

func (s *SqliteStore) DeleteProduct(id int64) error {
	query := `update product
	set 
		is_deleted = $1
	where
		id = $2
	`

	_, err := s.db.Exec(
		query,
		true,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

const file string = "products.db"

func NewSqliteStore() (*SqliteStore, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SqliteStore{
		db: db,
	}, nil
}

func (s *SqliteStore) Init() error {
	return s.createProductsTable()
}

func (s *SqliteStore) createProductsTable() error {
	fmt.Println("generating product table")
	query := `
	drop table product;
	create table product (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		currency TEXT NOT NULL DEFAULT "euro",
		in_stock INTEGER NOT NULL DEFAULT(0),
		is_deleted INTEGER NOT NULL DEFAULT(0)
	);
	CREATE unique INDEX if not exists unique_name ON product (name, is_deleted);
	`

	_, err := s.db.Exec(query)
	if err == nil {
		fmt.Println("products table created")
	}
	return err
}

func scanIntoProduct(rows *sql.Rows) (*Product, error) {
	p := new(Product)
	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Currency,
		&p.InStock,
		&p.IsDeleted,
	)
	return p, err
}
