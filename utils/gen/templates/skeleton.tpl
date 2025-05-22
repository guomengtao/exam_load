package skeleton

// {{.ControllerName}}Skeleton 骨架层
type {{.ControllerName}}Skeleton struct{}

// Hello 返回默认信息
func (s *{{.ControllerName}}Skeleton) Hello() string {
	return "{{.HelloMessage}}"
}