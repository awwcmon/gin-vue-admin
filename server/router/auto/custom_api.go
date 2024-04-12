package auto

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomApiRouter struct {
}

// InitCustomApiRouter 初始化 CustomApi 路由信息
func (s *CustomApiRouter) InitCustomApiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	customApiRouter := Router.Group("customApi").Use(middleware.OperationRecord())
	customApiRouterWithoutRecord := Router.Group("customApi")
	customApiRouterWithoutAuth := PublicRouter.Group("customApi")

	var customApiApi = v1.ApiGroupApp.AutoApiGroup.CustomApiApi
	{
		customApiRouter.POST("createCustomApi", customApiApi.CreateCustomApi)             // 新建CustomApi
		customApiRouter.DELETE("deleteCustomApi", customApiApi.DeleteCustomApi)           // 删除CustomApi
		customApiRouter.DELETE("deleteCustomApiByIds", customApiApi.DeleteCustomApiByIds) // 批量删除CustomApi
		customApiRouter.PUT("updateCustomApi", customApiApi.UpdateCustomApi)              // 更新CustomApi
		customApiRouter.POST("execute", customApiApi.Execute)                             //执行customApi
	}
	{
		customApiRouterWithoutRecord.GET("findCustomApi", customApiApi.FindCustomApi)       // 根据ID获取CustomApi
		customApiRouterWithoutRecord.GET("getCustomApiList", customApiApi.GetCustomApiList) // 获取CustomApi列表
	}
	{
		customApiRouterWithoutAuth.GET("getCustomApiPublic", customApiApi.GetCustomApiPublic) // 获取CustomApi列表
	}
}
