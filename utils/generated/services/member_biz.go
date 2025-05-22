package generated

// MemberBiz provides business logic layer for member
type MemberBiz struct {}

var MemberServiceInstance = NewMemberService()

func NewMemberBiz() *MemberBiz {
	return &MemberBiz{}
}

// ExampleBizMethod demonstrates calling service method from biz layer
func (b *MemberBiz) ExampleBizMethod() {
	// Calling service method from biz layer
	MemberServiceInstance.Example()
}
