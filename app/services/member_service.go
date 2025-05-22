package services

import (
	"fmt"

	gen "gin-go-test/utils/generated/services"
)

var MemberServiceInstance = gen.NewMemberService()

type MemberBiz struct {}

func NewMemberBiz() *MemberBiz {
	return &MemberBiz{}
}

func (b *MemberBiz) ExampleBizMethod() {
	fmt.Println("Calling service method from app biz layer")
	MemberServiceInstance.Example()
}
