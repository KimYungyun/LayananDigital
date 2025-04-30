package main

import (
	"database/sql"
	"html/template" //
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID         int
	Name       string
	Price      int
	CategoryID int
	Category   string
}

type PageData struct {
	Categories []Category
	Products   []Product
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/topup")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tambah-kategori", tambahKategoriHandler)
	http.HandleFunc("/tambah-produk", tambahProdukHandler)
	http.HandleFunc("/update-kategori", updateKategoriHandler)
	http.HandleFunc("/delete-kategori", deleteKategoriHandler)
	http.HandleFunc("/update-produk", updateProdukHandler)
	http.HandleFunc("/delete-produk", deleteProdukHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // ✅ Tambahkan error handling
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		rows.Scan(&c.ID, &c.Name)
		categories = append(categories, c)
	}

	prodRows, err := db.Query(`SELECT p.id, p.name, p.price, c.id, c.name FROM products p LEFT JOIN categories c ON p.category_id = c.id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer prodRows.Close()

	var products []Product
	for prodRows.Next() {
		var p Product
		prodRows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.Category)
		products = append(products, p)
	}

	tmpl.Execute(w, PageData{Categories: categories, Products: products})
}

func tambahKategoriHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateKategoriHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		_, err := db.Exec("UPDATE categories SET name = ? WHERE id = ?", name, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteKategoriHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		_, err := db.Exec("DELETE FROM categories WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func tambahProdukHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		price := r.FormValue("price")
		categoryID := r.FormValue("category_id")
		_, err := db.Exec("INSERT INTO products (name, price, category_id) VALUES (?, ?, ?)", name, price, categoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateProdukHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		price := r.FormValue("price")
		categoryID := r.FormValue("category_id")
		_, err := db.Exec("UPDATE products SET name = ?, price = ?, category_id = ? WHERE id = ?", name, price, categoryID, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteProdukHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
