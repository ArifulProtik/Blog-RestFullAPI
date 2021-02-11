package controller

import (
	"forum-api/utils"
	"io"
	"net/http"
	"os"
	"strings"
)

// Uploader Image Uploader
func Uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
		utils.ErrorWriter(w, "Please Select a Image", http.StatusNotAcceptable)
		return

	}
	filename := utils.UIDGen() + ".jpg"
	f, err := os.OpenFile("./cdn/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	url := "http://localhost:8888:/data/" + filename
	utils.SuccessWriter(w, url, http.StatusCreated)

}
