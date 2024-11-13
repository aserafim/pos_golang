package main

/*
Logs de operações:

// Modifique suas funções para que, a cada operação de inserção,
// atualização ou exclusão de produto, seja criado um registro em uma
// tabela de logs. Esse registro pode conter o ID do produto,
// a operação realizada e a data/hora.
*/

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Log struct {
	IDLog     string
	IDProduct string
	Operation string
	DateTime  string
}

func NewLog(_IDProduct string, _Operation string) *Log {
	return &Log{
		IDLog:     uuid.New().String(),
		IDProduct: _IDProduct,
		Operation: _Operation,
		DateTime:  time.Nanosecond.String(),
	}
}

func writeLog(db *sql.DB, log Log) {
	stmt, err := db.Prepare("insert into logs(id, )")
}

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
	if product.ID == "" {
		return errors.New("id cannot be empty")
	}
	if product.Name == "" {
		return errors.New("name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("price cannot be less than zero")
	}
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
	if product.ID == "" {
		return errors.New("id cannot be empty")
	}
	if product.Name == "" {
		return errors.New("name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("price cannot be less than zero")
	}
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// 4 - Atualização em massa:
// Crie uma função que permita atualizar o preço de todos
// os produtos que estejam dentro de um determinado intervalo
// de preço (ex.: aumentar o preço de todos os produtos que
// custam entre R$ 100 e R$ 500 em 10%).
func updateWithFilter(db *sql.DB, minPrice float64, maxPrice float64, pctDesc float64) error {
	stmt, err := db.Prepare("update products set price = ? where id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	var products []Product
	products, err = selectByPrice(db, minPrice, maxPrice)
	if err != nil {
		return err
	}

	for _, product := range products {
		newPrice := product.Price * (1 - pctDesc/100)
		if product.ID == "" {
			return errors.New("id cannot be empty")
		}
		if product.Name == "" {
			return errors.New("name cannot be empty")
		}

		if product.Price < 0 {
			return errors.New("price cannot be less than zero")
		}
		_, err = stmt.Exec(newPrice, product.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// 5 - Exclusão em lote:

// Implemente uma função que exclua todos os produtos
// cujo preço esteja acima de um valor especificado.
func deleteWithFilter(db *sql.DB, maxPrice float64) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	var products []Product
	products, err = selectByPrice(db, maxPrice, 5000000)
	if err != nil {
		return err
	}
	for _, product := range products {
		if product.ID == "" {
			return errors.New("id cannot be empty")
		}
		if product.Name == "" {
			return errors.New("name cannot be empty")
		}

		if product.Price < 0 {
			return errors.New("price cannot be less than zero")
		}
		_, err = stmt.Exec(product.ID)
		if err != nil {
			return err
		}
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

	if p.ID == "" {
		return nil, errors.New("id cannot be empty")
	}
	if p.Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if p.Price < 0 {
		return nil, errors.New("price cannot be less than zero")
	}

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
		if p.ID == "" {
			return nil, errors.New("id cannot be empty")
		}
		if p.Name == "" {
			return nil, errors.New("name cannot be empty")
		}

		if p.Price < 0 {
			return nil, errors.New("price cannot be less than zero")
		}
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
func selectByPrice(db *sql.DB, minPrice float64, maxPrice float64) ([]Product, error) {
	query := "select id, name, price from products where price >= " + strconv.FormatFloat(minPrice, 'f', -1, 64) + " and price <= " + strconv.FormatFloat(maxPrice, 'f', -1, 64)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		if p.ID == "" {
			return nil, errors.New("id cannot be empty")
		}
		if p.Name == "" {
			return nil, errors.New("name cannot be empty")
		}

		if p.Price < 0 {
			return nil, errors.New("price cannot be less than zero")
		}
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// var prod1 Product
	// prod1.Name = "Teste"
	// prod1.Price = 0

	// fmt.Println(insertProduct(db, &prod1))

	//db2, err := sql.Open("sqlite", "")
	//var products []Product
	//products, err = selectByPrice(db, 5000, 10000)

	// err = updateWithFilter(db, 5000, 9000, 0.95)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, prd := range products {
	// 	fmt.Println(prd)
	// }

	// prd1 := NewProduct("Notebook Asser Turbo 5", 3890.00)
	// prd2 := NewProduct("Notebook Asser Turbo 7", 5400.00)
	// prd3 := NewProduct("Tablet Orange Pad", 9000.00)
	// products = append(products, *prd1, *prd2, *prd3)
	// for _, valor := range products {
	// 	err := insertProduct(db, &valor)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// deleteWithFilter(db, 9000)

	db, err := sql.Open("sqlite", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
