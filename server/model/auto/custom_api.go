// 自动生成模板CustomApi
package auto

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CustomApi 结构体  CustomApi
type CustomApi struct {
	global.GVA_MODEL
	ApiName     string `json:"apiName" form:"apiName" gorm:"primarykey;column:api_name;comment:;" binding:"required"`  //api名称
	Api         string `json:"api" form:"api" gorm:"column:api;comment:;" binding:"required"`                          //api
	Desc        string `json:"desc" form:"desc" gorm:"column:desc;comment:;" binding:"required"`                       //描述
	Method      string `json:"method" form:"method" gorm:"column:method;comment:;" binding:"required"`                 //请求类型
	ParamConfig string `json:"paramConfig" form:"paramConfig" gorm:"column:param_config;comment:;" binding:"required"` //参数
}

// TableName CustomApi CustomApi自定义表名 CustomApi
func (CustomApi) TableName() string {
	return "custom_api"
}
