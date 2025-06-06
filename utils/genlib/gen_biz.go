package genlib

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type BizData struct {
	ModelName    string
	TableName    string
	Package      string
	BizName      string
	Imports      string
	HasList      bool
	RoutePath    string
	ImportModels string
	ReturnType   string
}

func GenerateBiz(table string, hasList bool) error {
	modelName := toCamelCase(table)
	routePath := strings.ToLower(table)
	data := BizData{
		ModelName:    modelName,
		TableName:    table,
		Package:      "biz",
		BizName:      modelName + "Biz",
		Imports:      "\"gin-go-test/app/models\"",
		HasList:      hasList,
		RoutePath:    routePath,
		ImportModels: "gin-go-test/app/models",
		ReturnType:   "[]*models." + modelName,
	}

	bizTplPath := filepath.Join("utils", "gen", "templates", "biz.tpl")
	bizSkeletonTplPath := filepath.Join("utils", "gen", "templates", "biz_skeleton.tpl")

	bizTpl, err := template.ParseFiles(bizTplPath)
	if err != nil {
		log.Printf("Failed to parse biz template: %v", err)
		return err
	}

	bizSkeletonTpl, err := template.ParseFiles(bizSkeletonTplPath)
	if err != nil {
		log.Printf("Failed to parse biz skeleton template: %v", err)
		return err
	}

	bizDir := filepath.Join("app", "biz")
	if err := os.MkdirAll(bizDir, 0755); err != nil {
		log.Printf("Failed to create biz dir: %v", err)
		return err
	}

	bizFile := filepath.Join(bizDir, table+"_biz.go")
	bizSkeletonDir := filepath.Join("utils", "generated", "biz")
	if err := os.MkdirAll(bizSkeletonDir, 0755); err != nil {
		log.Printf("Failed to create biz skeleton dir: %v", err)
		return err
	}
	bizSkeletonFile := filepath.Join(bizSkeletonDir, table+"_biz_skeleton.go")

	var bizBuf bytes.Buffer
	bizBuf.WriteString("// Code generated by gen.go. DO NOT EDIT.\n")
	if err := bizTpl.Execute(&bizBuf, data); err != nil {
		log.Printf("Failed to execute biz template: %v", err)
		return err
	}

	var bizSkeletonBuf bytes.Buffer
	bizSkeletonBuf.WriteString("// Code generated by gen.go. DO NOT EDIT.\n")
	if err := bizSkeletonTpl.Execute(&bizSkeletonBuf, data); err != nil {
		log.Printf("Failed to execute biz skeleton template: %v", err)
		return err
	}

	if err := os.WriteFile(bizFile, bizBuf.Bytes(), 0644); err != nil {
		log.Printf("Failed to write biz file: %v", err)
		return err
	}
	if err := os.WriteFile(bizSkeletonFile, bizSkeletonBuf.Bytes(), 0644); err != nil {
		log.Printf("Failed to write biz skeleton file: %v", err)
		return err
	}
	return nil
}

func GenerateBizSimple(table string) error {
	return GenerateBiz(table, false)
}

func toCamelCase(s string) string {
	b := make([]byte, 0, len(s))
	upperNext := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' || c == '-' {
			upperNext = true
			continue
		}
		if upperNext && c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		b = append(b, c)
		upperNext = false
	}
	return string(b)
}

func GenerateBizSkeleton(table string, hasList bool) error {
	modelName := toCamelCase(table)
	data := BizData{
		ModelName:    modelName,
		TableName:    table,
		Package:      "biz",
		BizName:      modelName + "Biz",
		Imports:      "\"gin-go-test/app/models\"",
		HasList:      hasList,
		ImportModels: "gin-go-test/app/models",
		ReturnType:   "[]*models." + modelName,
	}

	bizSkeletonTplPath := filepath.Join("utils", "gen", "templates", "biz_skeleton.tpl")

	bizSkeletonTpl, err := template.ParseFiles(bizSkeletonTplPath)
	if err != nil {
		log.Printf("Failed to parse biz skeleton template: %v", err)
		return err
	}

	bizSkeletonDir := filepath.Join("utils", "generated", "biz")
	if err := os.MkdirAll(bizSkeletonDir, 0755); err != nil {
		log.Printf("Failed to create biz skeleton dir: %v", err)
		return err
	}

	bizSkeletonFile := filepath.Join(bizSkeletonDir, table+"_biz_skeleton.go")

	var bizSkeletonBuf bytes.Buffer
	bizSkeletonBuf.WriteString("// Code generated by gen.go. DO NOT EDIT.\n")
	if err := bizSkeletonTpl.Execute(&bizSkeletonBuf, data); err != nil {
		log.Printf("Failed to execute biz skeleton template: %v", err)
		return err
	}

	if err := os.WriteFile(bizSkeletonFile, bizSkeletonBuf.Bytes(), 0644); err != nil {
		log.Printf("Failed to write biz skeleton file: %v", err)
		return err
	}

	return nil
}
