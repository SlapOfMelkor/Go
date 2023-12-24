package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Product struct {
	ID       int
	Name     string
	Category string
}

func database() {
	var err error
	db, err = sql.Open("mysql", "melkor:Melkor1993.@tcp(localhost)/marketdb")
	if err != nil {
		log.Fatal(err)
	} //AYNI URUN 2. DEFA EKLENMESIN
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS Products (
			id INT AUTO_INCREMENT PRIMARY KEY,  
			name VARCHAR(255) NOT NULL,
			category VARCHAR(255))`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS Sales (
			id INT AUTO_INCREMENT PRIMARY KEY,
			product_id INT,
			sale_date DATE,
			quantity INT,
			FOREIGN KEY (product_id) REFERENCES Products(id))`)
	if err != nil {
		log.Fatal(err)
	}
}
func list() []Product {
	rows, err := db.Query("SELECT id, name, category FROM Products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Category)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}

	return products
}

func addProduct() {
	var name, category string
	fmt.Print("Ürün Adı: ")
	fmt.Scan(&name)
	fmt.Print("Kategori: ")
	fmt.Scan(&category)
	_, err := db.Exec("INSERT INTO Products (name, category) VALUES (?, ?)", name, category)
	if err != nil {
		log.Fatal(err)
	}
}
func addSale() {
	products := list()
	if len(products) == 0 {
		fmt.Println("Lutfen Once Urun Ekleyin")
		return
	}
	fmt.Println("Urun Listesi")
	for i, product := range products {
		fmt.Printf("%d. %s (%s)\n", i+1, product.Name, product.Category)
	}
	var productChoice, quantity int
	fmt.Print("Urun Secin ")
	fmt.Scan(&productChoice)
	if productChoice < 1 || productChoice > len(products) {
		fmt.Println("Yanlis Giris Yaptiniz")
		return
	}
	fmt.Print("Satis Miktari Girin: ")
	fmt.Scan(&quantity)
	if quantity <= 0 {
		fmt.Println("Gecersiz Miktar")
		return
	}
	productID := products[productChoice-1].ID
	saleDate := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO Sales (product_id, sale_date, quantity) VALUES (?, ?, ?)", productID, saleDate, quantity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d adet %s urun satisi eklendi\n", quantity, products[productChoice-1].Name)
}

func showTopSale() {
	rows, err := db.Query(`
		SELECT Products.name, SUM(Sales.quantity) as total_quantity
		FROM Products
		JOIN Sales ON Products.id = Sales.product_id
		GROUP BY Products.id
		ORDER BY total_quantity DESC
		LIMIT 5
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("En Cok Satilan Urunler")
	for rows.Next() {
		var productName string
		var totalQuantity int
		err := rows.Scan(&productName, &totalQuantity)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d adet\n", productName, totalQuantity)
	}
}

func showCategorySale() {
	rows, err := db.Query(`
		SELECT Products.category, SUM(Sales.quantity) as total_quantity
		FROM Products
		JOIN Sales ON Products.id = Sales.product_id
		GROUP BY Products.category
		ORDER BY total_quantity DESC
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Kategoriye Gore Satis Istatistikleri")
	for rows.Next() {
		var category string
		var totalQuantity int
		err := rows.Scan(&category, &totalQuantity)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d adet\n", category, totalQuantity)
	}
}

func main() {
	database()

	for {
		fmt.Println("Ana Menü:")
		fmt.Println("1. Urun Girisi")
		fmt.Println("2. Satis Girisi")
		fmt.Println("3. En Cok Satilan Urunler")
		fmt.Println("4. Kategoriye Gore Satis Istatistikleri")
		fmt.Println("0. Cikis")
		var choice int
		fmt.Print("Seciminiz: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addProduct()
		case 2:
			addSale()
		case 3:
			showTopSale()
		case 4:
			showCategorySale()
		case 0:
			fmt.Println("Cikis Yapiliyor...")
			defer db.Close()
			return
		default:
			fmt.Println("Yanlis Secim Yaptiniz, Tekrar Deneyin")
		}
	}
}
