package product

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	connection *sql.DB
}

func NewRepository(connection *sql.DB) *Repository {
	return &Repository{
		connection: connection,
	}
}

// GetProducts
func (r *Repository) GetProducts() ([]Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := r.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []Product{}, err
	}

	var productList []Product
	var productObj Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (r *Repository) CreateProduct(product Product) (int, error) {

	var id int
	query, err := r.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (r *Repository) GetProductById(id_product int) (*Product, error) {

	query, err := r.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (r *Repository) DeleteProductByID(id int) error {

	query := "DELETE FROM product WHERE id = $1"

	result, err := r.connection.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *Repository) UpdateProduct(product Product) error {

	query, err := r.connection.Prepare(
		"UPDATE product " +
			"SET product_name = $1, price = $2 " +
			"WHERE id = $3",
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(
		product.Name,
		product.Price,
		product.ID,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()

	return nil
}
