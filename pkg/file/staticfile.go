package file

import (
	"fmt"
	"github.com/Talingan-Backend/utils"
	"io/ioutil"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./tmp"

func UploadFileHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		utils.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
		utils.WrapAPIError(w,r,"unable to parse form", http.StatusInternalServerError)
		return
	}

	// parse and validate file and post parameters
	file, fileHeader, err := r.FormFile("talingan-file-iot")
	if err != nil {
		utils.WrapAPIError(w,r,"invalid file form key", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Get and print out file size
	fileSize := fileHeader.Size
	fmt.Printf("File size (bytes): %v\n", fileSize)
	// validate file size
	if fileSize > maxUploadSize {
		utils.WrapAPIError(w,r,"max file size is 2 MB",http.StatusBadRequest)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		utils.WrapAPIError(w,r,"invalid file reading",http.StatusBadRequest)
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		utils.WrapAPIError(w,r,"invalid file type", http.StatusBadRequest)
		return
	}
	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		utils.WrapAPIError(w,r,"error read file type",http.StatusInternalServerError)
		return
	}
	newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
	fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		utils.WrapAPIError(w,r,"error writing file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		utils.WrapAPIError(w,r,"error writing file", http.StatusInternalServerError)
		return
	}
	utils.WrapAPISuccess(w,r,"success uploading file",http.StatusOK)
}



func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
