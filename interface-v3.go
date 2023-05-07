package main

import "fmt"

type bola interface {
	tendang() string
	oper() string
}

type tujuan struct {
	musuh string
	kawan string
}

// type tess struct {
// 	isi	string
// }

func (t tujuan) tendang() string {
	return t.musuh
}

func (t tujuan) oper() string {
	return t.kawan
}

func main() {
	var sepakBola bola
	sepakBola = tujuan{
		"musuh", "kawan",
	}
	fmt.Println("Tendang ke :: ", sepakBola.tendang())
	fmt.Println("Oper ke :: ", sepakBola.oper())
}
