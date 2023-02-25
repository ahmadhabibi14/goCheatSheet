// Random Tipe Data String
/*
Untuk menghasilkan data random string, ada banyak cara yang bisa digunakan,
salah satunya adalah dengan memafaatkan alfabet dan hasil random numerik.
*/

package main

import "math/rand"

var letters []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random string 5 karakter:", randomString(5))
}

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}