package main
import "fmt"

func main() {
	var point = 4

	switch {
	case point == 8 :
		fmt.Println("Perfect")
	case (point < 8) && (point > 3) :
		fmt.Println("Awesome")
		fallthrough
		/*
			Seperti yang kita sudah singgung di atas, bahwa switch pada Golang memiliki beberapa
			perbedaan dengan bahasa lain. Ketika sebuah case terpenuhi, pengecekkan kondisi tidak
			akan diteruskan ke case-case setelahnya.

			Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case
			selanjutnya.
		*/
	case point < 5 :
		fmt.Println("You need to learn more")
		/*
			Setelah pengecekkan case (point < 8) && (point > 3) selesai, akan dilanjut
			ke pengecekkan case point < 5 , karena ada fallthrough di situ.
		*/
	default :
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn more !!")
		}
	}
}