package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	db "github.com/DanialKassym/GoStorage/cmd/rest-server/internal/Database"
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

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}

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

	if contentType == "application/pdf" {
		Savefile(file, "/pdf/", w, *header)
	} else {
		Savefile(file, "/txt/", w, *header)
	}
	defer file.Close()

	buf.Reset()
	w.WriteHeader(http.StatusOK)
}

func Show (w http.ResponseWriter, r *http.Request) {
	str := "hello world"
	w.Write([]byte(str))
}

/*func Authorize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    var req AuthRequest
    err = json.Unmarshal(body, &req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	if authentication.ValidateJWT(req.Token){
		w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Request is authorized")
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(w, "Unauthorized")
    }
}

type AuthRequest struct {
    Token   string `json:"token"`
}*/
