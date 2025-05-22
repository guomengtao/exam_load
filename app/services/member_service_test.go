package services

import (
	"testing"

	gen "gin-go-test/utils/generated/services"
)

func TestMemberBiz_Example(t *testing.T) {
	biz := NewMemberBiz()
	biz.ExampleBizMethod()
}
