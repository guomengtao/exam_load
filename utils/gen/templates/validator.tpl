package validators

import (
    "github.com/go-playground/validator/v10"
    "fmt"
    "strings"
    "gin-go-test/app/models"
)

// Validate{{.StructName}} 用于校验 {{.StructName}} 字段
func Validate{{.StructName}}(item *models.{{.StructName}}) []string {
    var errors []string
    validate := validator.New()
    err := validate.Struct(item)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            errors = append(errors, fmt.Sprintf("Field '%s' failed on the '%s' tag", err.Field(), err.Tag()))
        }
    }
    return errors
} 