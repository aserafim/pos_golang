package main

import (
	"database/sql"
	"fmt"

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

func selectAllProducts(db *sql.DB) ([]Product, error) {
	// como não serão passados parâmetros
	// não corremos o risco de um sql injection
	// assim, não é necessário o uso do prepare
	rows, err := db.Query("select id, name, price from products")
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

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// desafio! remover registro
// func deleteProduct(db *sql.DB, id string) (bool, error) {
// 	stmt, err := db.Prepare("delete from products where id=?")
// 	if err != nil {
// 		return false, err
// 	}
// 	defer stmt.Close()
// 	var p Product
// 	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// desafio! Trazer mais de um registro do banco
// func selectProducts(db *sql.DB) ([]*Product, error) {
// 	stmt, err := db.Prepare("select id, name, price from products")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()
// 	var p Product
// 	var lista = make([]*Product, 0)

// 	for err != nil {
// 		err = stmt.QueryRow(stmt).Scan(&p.ID, &p.Name, &p.Price)
// 		if err != nil {
// 			return nil, err
// 		}
// 		lista = append(lista, &p)
// 	}
// 	return lista, err
// }

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// prod := NewProduct("PineApple Phone", 3500.0)
	// err = insertProduct(db, prod)
	// if err != nil {
	// 	panic(err)
	// }

	// prod.Name = "PineApple Phone 5"
	// err = updateProduct(db, prod)
	// if err != nil {
	// 	panic(err)
	// }

	// p, err := selectProduct(db, prod.ID)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf(p.Name)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Produto %s está no valor de %2.f\n", p.Name, p.Price)
	}

	id_user := "f8ffeaa6-62e6-489a-9127-b1a75adb90fd"

	_ = deleteProduct(db, id_user)

	products_dois, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products_dois {
		fmt.Printf("Produto %s está no valor de %2.f\n", p.Name, p.Price)
	}
}
