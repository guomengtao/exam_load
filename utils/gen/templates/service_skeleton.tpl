package service

type {{ .ServiceName }}ServiceSkeleton struct {}

func (s *{{ .ServiceName }}ServiceSkeleton) ExampleMethod() string {
	return "hello_skeleton"
}

func (s *{{ .ServiceName }}ServiceSkeleton) GetCount() string {
	return "skeleton count"
}