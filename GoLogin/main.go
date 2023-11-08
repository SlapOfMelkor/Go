package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Print("Giris yapmak icin 5 deneme hakkiniz var, lutfen kullanici adi ve sifrenizi giriniz\n\n")

	const sistemKul = "melkor"
	const sistemSif = "melkor"

	var userKul, userSif string

	file, _ := os.Create("log.txt")

	for i := 5; i > 0; i-- {

		fmt.Print("\nKullanıcı adı:")
		fmt.Scanf("%s\n", &userKul)
		fmt.Print("Sifre:")
		fmt.Scanf("%s\n", &userSif)
		myTime := time.Now()
		if userKul == sistemKul && userSif == sistemSif {
			io.WriteString(file, "\n"+userKul+"\n"+userSif+"\nBasarili\n"+myTime.Format("01-02-2006 15:04:05"))
			fmt.Print("Giris Basarili Hos Geldiniz")
			break
		} else {
			io.WriteString(file, "\n"+userKul+"\n"+userSif+"\nBasarisiz\n"+myTime.Format("01-02-2006 15:04:05"))
			fmt.Print("Giris Basarisiz Tekrar Deneyin")
		}
	}

	file.Close()
}
