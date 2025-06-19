package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Ошибка при закрытии файла:", err)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.IdentifyData(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := strings.ReplaceAll(time.Now().UTC().Format("2006-01-02T15-04-05Z"), ":", "-")
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}
	newFileName := fmt.Sprintf("converted-%s%s", filename, ext)

	outFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			log.Println("Ошибка при закрытии файла:", err)
		}
	}()

	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plane; charset=utf-8")
	_, _ = w.Write([]byte(result))

}
