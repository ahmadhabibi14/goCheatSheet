package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Ambil file dari form
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Tentukan lokasi penyimpanan file
		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, 0755)
		}

		// Buat file baru di lokasi penyimpanan
		filename := filepath.Join(uploadDir, header.Filename)
		newFile, err := os.Create(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Salin file yang diupload ke file baru
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Beri tahu pengguna bahwa file telah terupload
		fmt.Fprintf(w, "File %s telah terupload!", header.Filename)
	} else {
		// Tampilkan form untuk mengupload file
		w.Write([]byte(
			<form method="POST" enctype="multipart/form-data">
				Pilih file yang akan diupload: <input type="file" name="file"><br>
				<input type="submit" value="Upload">
			</form>
		))
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
