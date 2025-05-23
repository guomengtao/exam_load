package genlib

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"
)

func GenerateBizListContent(data BizData) (string, error) {
	listTplPath := filepath.Join("utils", "gen", "templates", "biz", "list.tpl")
	listTpl, err := template.ParseFiles(listTplPath)
	if err != nil {
		log.Printf("Failed to parse biz list template: %v", err)
		return "", err
	}
	var listBuf bytes.Buffer
	if err := listTpl.Execute(&listBuf, data); err != nil {
		log.Printf("Failed to execute biz list template: %v", err)
		return "", err
	}
	return listBuf.String(), nil
}