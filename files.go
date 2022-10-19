package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// SaveUploadedFile uploads the form file to specific dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	path := strings.TrimSuffix(dst, "/"+fileName)
	err = os.MkdirAll(path, 0666)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
