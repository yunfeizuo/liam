package route

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("unexpected request form", err)
		w.WriteHeader(400)
		w.Write([]byte("unexpected request form"))
		return
	}
	defer file.Close()

	name := uuid.NewV4().String()
	path := "./upload/" + name
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	log.Println("uploaded file", path)
	fmt.Fprintf(w, "%s", "/download/"+name)
}
