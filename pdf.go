package utils

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ledongthuc/pdf"
)

func ReadPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)

	defer func(err error) {
		if err != nil {
			log.Println("[PDF plugin] Close file Error:", err)
		}
	}(f.Close())

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ReadPdf2(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
	return "", nil
}
