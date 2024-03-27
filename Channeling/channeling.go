/*
Runtime' da deadlock alma durumları.
*/
package main

import "fmt"

func main() {
	chn := make(chan int, 2)
	chn <- 1
	chn <- 2
	/*
		Yukarıda channel oluşturulurken 2 elemanı olduğu belirtiliyor.
		Bu durumda aşağıdaki gibi kanala 3ncü bir eleman eklenmek istenirse aşağıdaki gibi çalışma zamanı hatası alınır.

		SORU : Rust ile benzer bir şey yazılabilir mi? Yazılıyorsa derleyici bu deadlock durumunu fark edip derlemeyi keser mi?

		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
				/workspaces/GoLang_201/Channeling/channeling.go:13 +0x58
		exit status 2
	*/

	// ch <- 3
	fmt.Println(<-chn)
	fmt.Println(<-chn)
}
