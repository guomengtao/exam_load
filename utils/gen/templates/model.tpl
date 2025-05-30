package models

{{ if .Imports }}import (
{{- range .Imports }}
	"{{ . }}"
{{- end }}
){{- end }}

// {{ .ModelName }} 数据模型  Key
type {{ .ModelName }} struct {
{{- range .Fields }}
	{{ .Name }} *{{ .Type }} `gorm:"column:{{ .Column }}" json:"{{ .JSON }}" validate:"{{ if and .IsRequired (not (and .IsPrimaryKey .IsAutoIncrement)) }}required,{{ end }}max=255"` 
{{- end }}
}