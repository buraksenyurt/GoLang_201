package models

import (
	"database/sql"
	"fmt"
	"log"

	// Kullanılmayan paketlerin save işleminde buradan silinmemesi için _ operatörü kullanılabilir

	_ "github.com/lib/pq" // go get github.com/lib/pq ile modülü eklemek gerekebilir
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "AdventureWorks"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float32
}

var db *sql.DB

/*
init metodları modül çalışma zamanına import ediliğinde otomatik olarak yürütülür.
Bu nedenle uygulama çalışmaya başladığında veri tabanı nesnesi otomatik olarak hazırlanmış olur.

Notlar:
- Küçük harfle başlayan fonksiynlar o modülün private fonksiyonlarıdır.
- Modüller içerisinde birden fazla dosyanın init fonksiyonu olabilir.
- Çağırılmayan modüllerin init fonksiyonu çalışmaz.
- Bir modül dosyasında (örneğin product.go) birden fazla init fonksiyonu da tanımlanabilir.
- SORU : Neden birden fazla init fonksiyonu yazabiliyoruz. Bunlar paralel işletilip modülün ön gerekiliklerinin daha çabuk yüklenmesi mümkün olabilir mi?
*/
func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disabled", host, port, user, password, dbname)
	// sqlDriver'ı bulmak için https://go.dev/wiki/SQLDrivers
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertProduct(data Product) (int, string) {
	var (
		id    int
		title string
	)
	db.QueryRow("INSERT INTO product(title,description,price) VALUES($1,$2,$3) RETURNING id,title", data.Title, data.Description, data.Price).Scan(&id, &title)
	fmt.Println("Inserted row id ", id)
	return id, title
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE product SET (title=$2,description=$3,price=$4) WHERE id=$1", data.Id, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowAffacted, err := result.RowsAffected()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found")
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Updated row count is ", rowAffacted)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM product ORDER BY 1 DESC")
	defer rows.Close()

	if err != nil {
		// if err == sql.ErrNoRows {
		// 	fmt.Println("No records found")
		// 	return
		// }
		log.Fatal(err)
	}

	var products []*Product
	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	if length := len(products); length == 0 {
		fmt.Printf("No records found")
		return
	}

	for _, p := range products {
		fmt.Printf("%d - %s - %s, %2.f\n", p.Id, p.Title, p.Description, p.Price)
	}
}

func GetProductById(id int) {
	var title string
	err := db.QueryRow("SELECT title FROM product WHERE id=$1", id).Scan(&title)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No product found")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Title of the product is %s\n", title)
	}
}
