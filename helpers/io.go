package helpers

import (
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn((len(letters)))]
	}
	return string(b)
}

func ImageIsJpgOrPng(header *multipart.FileHeader) bool {
	ext := filepath.Ext(header.Filename)
	log.Println(ext)
	return ext == ".png" || ext == ".jpg"
}

func UploadImage(file multipart.File, header *multipart.FileHeader) string {
	dir := "assets/image"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("creating dir : %s", err.Error())
		}
	}
	imageName := randSeq(20) + filepath.Ext(header.Filename)

	fileLocation := filepath.Join(dir, imageName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		log.Println(err.Error())
		return ""
	}

	return fileLocation
}

func RemoveFile(r *http.Request, URL string) error {
	fileLoc := "." + strings.Trim(URL, r.Host)

	log.Println(fileLoc)
	return os.Remove(fileLoc)
}

func SaveToCSV(in interface{}) (string, error) {

	dir := "assets/csv"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("creating dir : %s", err.Error())
		}
	}
	fileName := randSeq(20) + ".csv"

	fileLocation := filepath.Join(dir, fileName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer targetFile.Close()

	err = gocsv.MarshalFile(in, targetFile)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return fileLocation, nil
}
