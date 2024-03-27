# GoLang_201

DT içerisinde düzenlenen GoLang eğitimine ait notların tutulduğu repodur.

## Hello World

İlk hello world uygulamasını aşağıdaki adımlarla oluşturdum.

```bash
# proje için bir klasör açtım
mkdir HelloWorld
cd HelloWorld

# bir go modülü başlattım
go mod init example/HelloWorld
# main fonksiyonunu içerecek go dosyasını oluşturdum
touch helloworld.go

# dosya içeriği hazırladım
# çalıştırmak için aşağıdaki komutu kullandım

go run .
```

İlk örneğin kod içeriği ise aşağıdaki gibi.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Konular

Eğitim sırasında işlediğimiz konular,

- Go Routine kavramı _(GoLang 101 eğitiminde değinilen konu üzerinden geçildi. Basit go çağrımı, Wait Group ve Channel kullanımları ile Deadlock durumları ele alındı)_
- DB operasyonları _(PostgreSQL üzerinden basit veritabanı işlemleri ele alındı. Insert, Update, Delete, Select...)_
- File I/O İşlemleri _(Dosya açma, okuma ve yazma işlemlerine değiniliyor)_
- JSON serileştirme işlemleri _(Basit JSON serileştirme işlemleri yanında merkez bankasından döviz kurlarını çekme örneği yapıldı)_
- API geliştirme _(Temel HTTP routing işlemleri ve gorilla mux kullanımına değinildi)_
- Unit Test _(En basit anlamda birim test nasıl yazılır konusuna değinildi)_ Bu projede utility_test dışında weather klasöründe yer alan weather_service_test isimli ikinci bir birim test daha bulunuyor. Bir go projesindeki tüm testleri çalıştırmak için root klasördeyken aşağıdaki terminal komutu kullanılabilir.

```bash
# Projenin alt klasörlerindeki testler de dahil tüm testlerin çalıştırılması için
# v=verbose anlamındadır.
go test ./... -v
```

## Yardımcı Bilgiler

- Go veri tabanı sürücüleri ile ilgili olarak [bu adresten](https://go.dev/wiki/SQLDrivers) yararlanılabilir.
- Go ile web uygulamalarının yazılmasına ait [şu sitede](https://gowebexamples.com) güzel örnekler var.