package cheatsheet

import (
	"fmt"
	"testing"
)

type studentP struct {
	name  string
	grade int
}

func TestStructPointer(t *testing.T) {
	/*
		Objek hasil cetakan struct bisa diambil nilai pointer-nya, dan bisa disimpan
		pada variabel objek yang bertipe struct pointer
	*/
	var s1 = studentP{
		name:  "Wick",
		grade: 2,
	}
	var s2 *studentP = &s1
	fmt.Println("Student 1, name :", s1.name)
	fmt.Println("Student 4, name :", s2.name)

	s2.name = "Ethan"
	fmt.Println("Student 1, name :", s1.name)
	fmt.Println("Student 4, name :", s2.name)

	/*
		s2 adalah variabel pointer hasil cetakan struct student . s2 menampung nilai referensi
		s1 , mengakibatkan setiap perubahan pada property variabel tersebut, akan juga
		berpengaruh pada variabel objek s1 .

		Meskipun s2 bukan variabel asli, property nya tetap bisa diakses seperti biasa. Inilah
		keunikan variabel objek pointer, tanpa perlu di-dereferensi nilai asli property tetap bisa
		diakses. Pengisian nilai pada property tersebut juga bisa langsung menggunakan nilai asli,
		contohnya seperti s2.name = "ethan".
	*/
}
