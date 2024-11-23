package cheatsheet

import (
	"fmt"
	"testing"
)

func TestArraySliceV2(t *testing.T) {
	/*
		Slice merupakan tipe reference. Artinya jika ada slice baru yang terbentuk dari slice lama,
		maka elemen slice baru memiliki referensi yang sama dengan elemen slice lama.
		Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada
		elemen slice lama yang memiliki referensi yang sama
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]
	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	/*
		Variabel aFruits , bFruits merupakan slice baru yang terbentuk dari variabel fruits.
		Dengan menggunakan dua slice baru tersebut, diciptakan lagi slice lainnya, yaitu aaFruits,
		dan baFruits . Kelima slice tersebut ditampilkan nilainya. Selanjutnya, nilai
		dari baFruits[0] diubah, dan 5 slice tadi ditampilkan lagi. Hasilnya akan ada
		banyak slice yang elemennya ikut berubah. Yaitu elemen-elemen yang
		referensi-nya sama dengan referensi elemen baFruits[0]
	*/

	/*
		Bisa dilihat pada output di atas, elemen yang sebelumnya bernilai "grape" pada variabel
		fruits , aFruits , bFruits , aaFruits , dan baFruits ; kesemuanya berubah menjadi
		"pinnaple" , karena memiliki referensi yang sama.
	*/
}
