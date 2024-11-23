package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArrayMultiDimension(t *testing.T) {
	/*
		Array multidimensi adalah array yang tiap elemennya juga
		berupa array (dan bisa seterusnya, tergantung jumlah dimensinya).

		Khusus untuk array yang merupakan sub dimensi atau elemen,
		boleh tidak dituliskan jumlah datanya. Contohnya bisa dilihat
		pada deklarasi variabel numbers2 di kode berikut.
	*/
	var numbers1 = [2][3]int{
		[3]int{
			3, 2, 2,
		},
		[3]int{
			3, 4, 5,
		},
	}

	var numbers2 = [2][3]int{
		{3, 2, 3},
		{3, 4, 5},
	}

	// Kedua array di atas adalah sama nilainya.
	fmt.Println("Numbers1", numbers1)
	fmt.Println("Numbers2", numbers2)
}
