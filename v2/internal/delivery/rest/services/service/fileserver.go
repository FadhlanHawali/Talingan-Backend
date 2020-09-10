package service

import (
	"fmt"
	"github.com/Talingan-Backend/v2/internal/helper"
	"github.com/h2non/filetype"
	"io/ioutil"
	"math/rand"
	_ "mime"
	"net/http"
	"os"
	"path/filepath"
)

func (s *ServicesService) UploadFileHandler(w http.ResponseWriter, r *http.Request){
	const maxUploadSize = 2 * 1024 * 1024 // 2 mb
	const uploadPath = "./tmp"

	if r.Method != "POST" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
		helper.WrapAPIError(w,r,"unable to parse form", http.StatusInternalServerError)
		return
	}

	// parse and validate file and post parameters
	file, fileHeader, err := r.FormFile("talingan-file-iot")
	if err != nil {
		helper.WrapAPIError(w,r,"invalid file form key", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Get and print out file size
	fileSize := fileHeader.Size
	fmt.Printf("File size (bytes): %v\n", fileSize)
	// validate file size
	if fileSize > maxUploadSize {
		helper.WrapAPIError(w,r,"max file size is 2 MB",http.StatusBadRequest)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)

	kind, _ := filetype.Match(fileBytes)
	if kind == filetype.Unknown {
		helper.WrapAPIError(w,r,"unknown file type" + kind.MIME.Value, http.StatusBadRequest)
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
	if err != nil {
		helper.WrapAPIError(w,r,"invalid file reading",http.StatusBadRequest)
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	switch kind.MIME.Value {
	case "audio/mpeg":
		break
	default:
		helper.WrapAPIError(w,r,"wrong file type : " + kind.MIME.Value, http.StatusBadRequest)
		return
	}
	fileName := helper.IdGenerator()
	newPath := filepath.Join(uploadPath, fileName+"."+kind.Extension)
	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		helper.WrapAPIError(w,r,"error writing file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		helper.WrapAPIError(w,r,"error writing file", http.StatusInternalServerError)
		return
	}
	helper.WrapAPISuccess(w,r,"success uploading file",http.StatusOK)
}



func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}