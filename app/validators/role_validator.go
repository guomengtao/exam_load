package validators

import (
	"gin-go-test/app/models"
	"github.com/go-playground/validator/v10"
)

// RoleValidator 角色验证器
type RoleValidator struct {
	validate *validator.Validate
}

// NewRoleValidator 创建角色验证器
func NewRoleValidator() *RoleValidator {
	return &RoleValidator{
		validate: validator.New(),
	}
}

// Validate 验证单个对象
func (v *RoleValidator) Validate(role *models.Role) error {
	return v.validate.Struct(role)
}

// ValidateBatch 批量验证对象
func (v *RoleValidator) ValidateBatch(roles []*models.Role) []error {
	var errors []error
	for _, role := range roles {
		if err := v.Validate(role); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// ValidateCreate 验证创建请求
func (v *RoleValidator) ValidateCreate(role *models.Role) error {
	return v.Validate(role)
}

// ValidateUpdate 验证更新请求
func (v *RoleValidator) ValidateUpdate(role *models.Role) error {
	return v.Validate(role)
}

// ValidateBatchCreate 验证批量创建请求
func (v *RoleValidator) ValidateBatchCreate(roles []*models.Role) []error {
	return v.ValidateBatch(roles)
}

// ValidateBatchUpdate 验证批量更新请求
func (v *RoleValidator) ValidateBatchUpdate(roles []*models.Role) []error {
	return v.ValidateBatch(roles)
}
