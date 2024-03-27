/*
Bu ilk örnekte Go Routine kullanımı hatırlandı.
Senkrondan asenkrona doğru geçiş yapıldı.
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	var waitGroup sync.WaitGroup

	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	for _, url := range urls {
		// checkUrl(url) // senkron versiyon
		// go checkUrl(url) // fonksiyon asenkron olarak çağrıldı

		// Asenkron operasyonları waitable olarak
		waitGroup.Add(1)
		go func(url string) {
			defer waitGroup.Done()
			checkUrl(url)
		}(url)
	}

	waitGroup.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Total execution time is :%s\n", elapsedTime)
}

// Senkron url kontrolü yapan versiyon
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " isn't available")
	} else {
		fmt.Println(url, " is up and running")
	}
}
