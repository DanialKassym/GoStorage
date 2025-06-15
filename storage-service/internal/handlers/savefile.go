package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Savefile(file multipart.File, ext string, w http.ResponseWriter, header multipart.FileHeader) {
	id := generateUniqueID()
	filename := strconv.FormatInt(id, 10) + header.Filename
	dstPath := filepath.Join("/tmp", ext, filename)

	fileSize, _ := file.Seek(0, io.SeekEnd)
	file.Seek(0, io.SeekStart) // Reset file pointer to the beginning as it was moved till the end after being copied to buffer
	fmt.Println("File size:", fileSize)

	if fileSize == 0 {
		http.Error(w, "Empty file", http.StatusBadRequest)
		return
	}
	
	// Ensure the directory exists before writing the file
	err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to create directory", http.StatusInternalServerError)
		fmt.Println("Directory creation error:", err)
		return
	}

	// Create the destination file
	out, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}

	//copy the multipart to the file
	wr, err := io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer out.Close()

	fmt.Println(wr)
}

func generateUniqueID() int64 {
	return time.Now().UnixNano() + int64(rand.Intn(100))
}
