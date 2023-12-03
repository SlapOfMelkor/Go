package main

import "fmt"

func main() {

	var metin string = "dqwdqwd"
	var sayi = 12
	var sayi2 int = 15

	fmt.Println(metin)
	fmt.Println(sayi)
	fmt.Println(sayi2)
	fmt.Println(15 + (15 * 12))
	var deger bool = true

	fmt.Println(deger)

	var x float32 = 3.2

	fmt.Printf("%T", x)
	fmt.Println("")
	var text1 string = "aaaa"
	var text2 string = "bbb"

	var durum = text1 == text2
	fmt.Println(durum)
	fmt.Println(!durum)

}
