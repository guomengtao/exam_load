package biz

import (
	"errors"                 // 导入errors包用于自定义错误
	"gin-go-test/app/models" // 修正为正确的模块路径
	"gin-go-test/utils"
	"github.com/stretchr/testify/assert" // 导入assert用于断言
	"gorm.io/driver/sqlite"              // 导入sqlite驱动用于内存数据库
	"gorm.io/gorm"                       // 导入gorm用于数据库操作
	"testing"                            // 导入testing包用于单元测试
)

// setupTestDB 用于初始化内存数据库和迁移Role表结构
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) // 使用sqlite内存数据库
	assert.NoError(t, err)                                        // 断言数据库打开无错误
	err = db.AutoMigrate(&models.Role{})                          // 自动迁移Role表结构
	assert.NoError(t, err)                                        // 断言迁移无错误
	return db                                                     // 返回数据库连接
}

// TestRoleBiz_BatchCreate_Success 测试批量创建成功时事务提交
func TestRoleBiz_BatchCreate_Success(t *testing.T) {
	db := setupTestDB(t)  // 初始化测试数据库
	biz := NewRoleBiz(db) // 创建RoleBiz实例
	roles := []models.Role{
		{Name: "admin"}, // 添加一个角色admin
		{Name: "user"},  // 添加一个角色user
	}
	err := biz.BatchCreate(roles) // 调用批量创建方法
	assert.NoError(t, err)        // 断言无错误

	var count int64
	db.Model(&models.Role{}).Count(&count) // 统计Role表中的记录数
	assert.Equal(t, int64(2), count)       // 断言有2条记录
}

// TestRoleBiz_BatchCreate_TransactionRollback 测试批量创建失败时事务回滚
func TestRoleBiz_BatchCreate_TransactionRollback(t *testing.T) {
	db := setupTestDB(t)  // 初始化测试数据库
	biz := NewRoleBiz(db) // 创建RoleBiz实例
	roles := []models.Role{
		{Name: "admin"}, // 添加一个角色admin
		{Name: ""},      // 添加一个空Name角色，应该触发回滚
	}
	err := biz.BatchCreate(roles) // 调用批量创建方法
	assert.Error(t, err)          // 断言有错误发生

	var count int64
	db.Model(&models.Role{}).Count(&count) // 统计Role表中的记录数
	assert.Equal(t, int64(0), count)       // 断言没有记录，说明事务回滚
}

// RoleBiz 业务层结构体，包含数据库连接
// 实际项目中应由生成器生成
// 这里只做测试用例演示

type RoleBiz struct {
	db *gorm.DB // 数据库连接
}

// NewRoleBiz 创建RoleBiz实例
func NewRoleBiz(db *gorm.DB) *RoleBiz {
	return &RoleBiz{db: db}
}

// BatchCreate 批量创建角色，带事务和业务判断
func (b *RoleBiz) BatchCreate(roles []models.Role) error {
	// 使用事务包裹批量创建逻辑
	return b.db.Transaction(func(tx *gorm.DB) error {
		for _, role := range roles { // 遍历每个角色
			if role.Name == "" { // 业务判断：Name不能为空
				return errors.New("name is required") // 返回错误，事务将回滚
			}
			// 实际应调用service层的Create方法，这里直接用GORM演示
			if err := tx.Create(&role).Error; err != nil { // 创建角色
				return err // 有错误则回滚事务
			}
		}
		return nil // 全部成功则提交事务
	})
}

func setupRoleBiz() *RoleBiz {
	// 初始化 GORM 数据库连接
	utils.InitGorm()

	// 创建 RoleService 实例
	service := services.NewRoleService(utils.GormDB)

	// 创建 RoleBiz 实例
	return NewRoleBiz(utils.GormDB)
}

func TestRoleBiz_GetCount(t *testing.T) {
	biz := setupRoleBiz()

	// 调用 GetCount 方法
	count, err := biz.GetCount()
	if err != nil {
		t.Fatalf("调用 GetCount 失败: %v", err)
	}

	t.Logf("角色总数: %d", count)
}

func TestRoleBiz_List(t *testing.T) {
	biz := setupRoleBiz()

	// 调用 List 方法，分页参数：第 1 页，每页 10 条
	items, total, err := biz.List(1, 10)
	if err != nil {
		t.Fatalf("调用 List 失败: %v", err)
	}

	t.Logf("分页获取角色，共 %d 条记录，当前页返回 %d 条", total, len(items))
}

func TestRoleBiz_BatchCreate(t *testing.T) {
	// 初始化 biz 层
	biz := NewRoleBiz(utils.GormDB)

	// 测试用例1：正常创建
	t.Run("正常创建", func(t *testing.T) {
		items := []*models.Role{
			{Name: "测试角色1", Desc: "描述1"},
			{Name: "测试角色2", Desc: "描述2"},
		}
		err := biz.BatchCreate(items)
		assert.NoError(t, err)
	})

	// 测试用例2：超过30条记录
	t.Run("超过30条记录", func(t *testing.T) {
		items := make([]*models.Role, 31)
		for i := range items {
			items[i] = &models.Role{Name: "测试角色", Desc: "描述"}
		}
		err := biz.BatchCreate(items)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "maximum 30 records allowed")
	})

	// 测试用例3：空名称
	t.Run("空名称", func(t *testing.T) {
		items := []*models.Role{
			{Name: "", Desc: "描述"},
		}
		err := biz.BatchCreate(items)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})
}

func TestRoleBiz_List(t *testing.T) {
	// 初始化 biz 层
	biz := NewRoleBiz(utils.GormDB)

	// 测试用例1：基本分页
	t.Run("基本分页", func(t *testing.T) {
		params := SearchParams{
			Page:     1,
			PageSize: 10,
		}
		items, total, err := biz.List(params)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, total, int64(0))
		assert.LessOrEqual(t, len(items), 10)
	})

	// 测试用例2：带关键字搜索
	t.Run("带关键字搜索", func(t *testing.T) {
		params := SearchParams{
			Page:     1,
			PageSize: 10,
			Keyword:  "测试",
		}
		items, total, err := biz.List(params)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, total, int64(0))
	})

	// 测试用例3：带过滤条件
	t.Run("带过滤条件", func(t *testing.T) {
		params := SearchParams{
			Page:     1,
			PageSize: 10,
			Filters: map[string]string{
				"name": "测试角色1",
			},
		}
		items, total, err := biz.List(params)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, total, int64(0))
	})
}

func TestRoleBiz_BatchUpdate(t *testing.T) {
	// 初始化 biz 层
	biz := NewRoleBiz(utils.GormDB)

	// 测试用例1：正常更新
	t.Run("正常更新", func(t *testing.T) {
		// 先创建测试数据
		items := []*models.Role{
			{Name: "更新测试1", Desc: "描述1"},
		}
		err := biz.BatchCreate(items)
		assert.NoError(t, err)

		// 更新数据
		items[0].Desc = "更新后的描述"
		err = biz.BatchUpdate(items)
		assert.NoError(t, err)
	})

	// 测试用例2：ID不存在
	t.Run("ID不存在", func(t *testing.T) {
		items := []*models.Role{
			{Name: "测试角色", Desc: "描述"},
		}
		err := biz.BatchUpdate(items)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required for update")
	})
}

func TestRoleBiz_BatchDelete(t *testing.T) {
	// 初始化 biz 层
	biz := NewRoleBiz(utils.GormDB)

	// 测试用例1：正常删除
	t.Run("正常删除", func(t *testing.T) {
		// 先创建测试数据
		items := []*models.Role{
			{Name: "删除测试1", Desc: "描述1"},
		}
		err := biz.BatchCreate(items)
		assert.NoError(t, err)

		// 删除数据
		err = biz.BatchDelete([]uint{items[0].ID})
		assert.NoError(t, err)
	})

	// 测试用例2：无效ID
	t.Run("无效ID", func(t *testing.T) {
		err := biz.BatchDelete([]uint{0})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be greater than 0")
	})
}
