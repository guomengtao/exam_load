package models

{{ if .Imports }}import (
{{- range .Imports }}
	"{{ . }}"
{{- end }}
){{- end }}

// {{ .ModelName }} 数据模型
type {{ .ModelName }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} `gorm:"column:{{ .Column }}" json:"{{ .JSON }}"` 
{{- end }}
}