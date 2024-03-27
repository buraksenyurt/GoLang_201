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

İlk örneğin kod içeriği de şöyleyi.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```