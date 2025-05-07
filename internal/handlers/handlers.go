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

// Used for db connection testing
func RetriveUsers(w http.ResponseWriter, r *http.Request) {

	ret := db.GetAllUsers()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

// supports pdf and text files
func UploadObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseMultipartForm(512 << 20)
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file: ", http.StatusBadRequest)
		return
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name is: %s\n", name[0])

	var buf bytes.Buffer
	n, _ := io.Copy(&buf, file)
	fmt.Printf("Read %d bytes\n", n)

	// Reading the first 512 bytes to detect content type
	contentType := http.DetectContentType(buf.Bytes()[:512])
	fmt.Println("Detected content type:", contentType)

	if contentType != "application/pdf" && contentType != "application/octet-stream" {
		http.Error(w, "Unsupported content type: ", http.StatusBadRequest)
		return
	}
	contents := buf.String()
	fmt.Println(contents)
	f, _ := os.Create("/tmp/data")
	defer f.Close()
	n2, _ := f.Write([]byte(contents))
	fmt.Printf("wrote %d bytes\n", n2)
	buf.Reset()
}
