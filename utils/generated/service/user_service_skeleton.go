package service

type UserServiceSkeleton struct {}

func (s *UserServiceSkeleton) ExampleMethod() string {
	return "hello_skeleton"
}

func (s *UserServiceSkeleton) GetCount() string {
	return "skeleton count"
}