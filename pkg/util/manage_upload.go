package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ManageImageUpload(r *http.Request) (string, error) {
	file, handler, err := r.FormFile("display_pic")
	if err != nil {
		return "", nil
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

	uploadPath := filepath.Join("uploads", fileName)
	dest, err := os.Create(uploadPath)
	if err != nil {
		return "", fmt.Errorf("error creating upload file : %v", err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return "", fmt.Errorf("Error copying file content: %v", err)
	}

	return uploadPath, nil
}
