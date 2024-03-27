/*
Bu ilk örnekte Go Routine kullanımı hatırlandı.
Senkrondan asenkrona doğru geçiş yapıldı.

go kullanımı
Wait Group kullanımı
channel kullanımı
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()

	//var waitGroup sync.WaitGroup

	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	chn := make(chan urlStatus)

	for _, url := range urls {
		go checkUrl(url, chn) // Kanal yardımıyla asenkron çağrı

		// checkUrl(url) // senkron versiyon
		// go checkUrl(url) // fonksiyon asenkron olarak çağrıldı

		// Asenkron operasyonları waitable olarak
		// waitGroup.Add(1)
		// go func(url string) {
		// 	defer waitGroup.Done()
		// 	checkUrl(url)
		// }(url)
	}

	//waitGroup.Wait()

	for i := 0; i < len(urls); i++ {
		urlS := <-chn
		if urlS.status {
			fmt.Println(urlS.url, " is up running")
		} else {
			fmt.Println(urlS.url, " is dead")
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Total execution time is :%s\n", elapsedTime)
}

// Senkron url kontrolü yapan versiyon
func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- urlStatus{url, false}
		//fmt.Println(url, " isn't available")
	} else {
		c <- urlStatus{url, true}
		//fmt.Println(url, " is up and running")
	}
}

type urlStatus struct {
	url    string
	status bool
}
