package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// En hızlı haliyle dosya yazma (Fast Writing)
	var err = os.WriteFile("data/products.txt", []byte("Product informations"), os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	// Dosya oluşturma ve içerisine basit bir metin yerleştirme
	newFile, err := os.Create("Games.dat")
	if err != nil {
		log.Fatal(err)
	}
	newFile.WriteString("Games data")

	// Dosya ile ilgili bilgileri alma
	fileInfo, err := os.Stat("Games.dat")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
		log.Fatal(err)
	}
	fmt.Println("File name ", fileInfo.Name())
	fmt.Println("File size ", fileInfo.Size())
	fmt.Println("File permissions ", fileInfo.Mode())
	fmt.Println("File last modified date ", fileInfo.ModTime())
	fmt.Println(fileInfo)

	newFile.Close()

	// Dosya taşıma
	orgPath := "Games.dat"
	trgPath := "data/games.dat"
	err = os.Rename(orgPath, trgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Dosya açma ve ek yapma
	file, err := os.OpenFile("data/games.dat", os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content := []byte("Pac Man 1984\n")
	len, err := file.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes write\n", len)
}
