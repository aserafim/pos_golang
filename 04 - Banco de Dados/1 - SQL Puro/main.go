package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	// Protege contra sql injection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// O _ abaixo, se utilizado como res, por exemplo
	// pode trazer algumas infos sobre a execução
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// 2 - Paginação de resultados:
// Modifique a função selectAllProducts para retornar um número
// limitado de produtos por vez (ex.: 10 produtos), implementando paginação.
func selectAllProducts(db *sql.DB, limit int) ([]Product, error) {
	// como não serão passados parâmetros
	// não corremos o risco de um sql injection
	// assim, não é necessário o uso do prepare
	query := "select id, name, price from products"
	if limit > 0 {
		query = "select id, name, price from products limit " + strconv.Itoa(limit)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// 1 - Inserir múltiplos produtos de uma vez:
// Crie uma função que receba um slice de Product e
// insira todos os produtos de uma só vez no banco
// de dados usando uma transação.
func IsertMany(db *sql.DB, products []*Product) error {
	for _, prd := range products {
		err := insertProduct(db, prd)
		if err != nil {
			return nil
		}
	}
	return nil
}

// 3 - Filtragem de produtos:
// Adicione um filtro por preço e nome. 
// Crie uma função que selecione produtos 
// com base em um intervalo de preços e/ou uma parte do nome.


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var products []Product

	products, err = selectAllProducts(db, 0)
	if err != nil {
		panic(err)
	}

	for _, prd := range products {
		fmt.Println(prd)
	}

	// prd1 := NewProduct("Notebook Asser Turbo 5", 3890.00)
	// prd2 := NewProduct("Notebook Asser Turbo 7", 5400.00)
	// prd3 := NewProduct("Tablet Orange Pad", 9000.00)

	// products = append(products, prd1, prd2, prd3)

	// for _, valor := range products {
	// 	err := insertProduct(db, valor)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
