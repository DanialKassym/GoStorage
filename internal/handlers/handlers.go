package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	db "github.com/DanialKassym/GoStorage/internal/Database"
)

func RetriveUsers(w http.ResponseWriter, r *http.Request) {

	ret := db.GetAllUsers()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

func UploadObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseMultipartForm(512 << 20)
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
	var buf bytes.Buffer
	fmt.Println("Before FormFile")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+ err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name is: %s\n", name[0])

	n, _ := io.Copy(&buf, file)
	fmt.Printf("Read %d bytes\n", n)
	contents := buf.String()
	fmt.Println(contents)

	f, _ := os.Create("/tmp/data")
	defer f.Close()
	n2, _ := f.Write([]byte(contents))
	fmt.Printf("wrote %d bytes\n", n2)
	buf.Reset()
}
