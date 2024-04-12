package auto

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/auto"
	autoReq "github.com/flipped-aurora/gin-vue-admin/server/model/auto/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomApiApi struct {
}

var customApiService = service.ServiceGroupApp.AutoServiceGroup.CustomApiService

// CreateCustomApi 创建CustomApi
// @Tags CustomApi
// @Summary 创建CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body auto.CustomApi true "创建CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customApi/createCustomApi [post]
func (customApiApi *CustomApiApi) CreateCustomApi(c *gin.Context) {
	var customApi auto.CustomApi
	err := c.ShouldBindJSON(&customApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := customApiService.CreateCustomApi(&customApi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomApi 删除CustomApi
// @Tags CustomApi
// @Summary 删除CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body auto.CustomApi true "删除CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customApi/deleteCustomApi [delete]
func (customApiApi *CustomApiApi) DeleteCustomApi(c *gin.Context) {
	ID := c.Query("ID")
	if err := customApiService.DeleteCustomApi(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCustomApiByIds 批量删除CustomApi
// @Tags CustomApi
// @Summary 批量删除CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /customApi/deleteCustomApiByIds [delete]
func (customApiApi *CustomApiApi) DeleteCustomApiByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := customApiService.DeleteCustomApiByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCustomApi 更新CustomApi
// @Tags CustomApi
// @Summary 更新CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body auto.CustomApi true "更新CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customApi/updateCustomApi [put]
func (customApiApi *CustomApiApi) UpdateCustomApi(c *gin.Context) {
	var customApi auto.CustomApi
	err := c.ShouldBindJSON(&customApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := customApiService.UpdateCustomApi(customApi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCustomApi 用id查询CustomApi
// @Tags CustomApi
// @Summary 用id查询CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query auto.CustomApi true "用id查询CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customApi/findCustomApi [get]
func (customApiApi *CustomApiApi) FindCustomApi(c *gin.Context) {
	ID := c.Query("ID")
	if recustomApi, err := customApiService.GetCustomApi(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recustomApi": recustomApi}, c)
	}
}

// GetCustomApiList 分页获取CustomApi列表
// @Tags CustomApi
// @Summary 分页获取CustomApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autoReq.CustomApiSearch true "分页获取CustomApi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customApi/getCustomApiList [get]
func (customApiApi *CustomApiApi) GetCustomApiList(c *gin.Context) {
	var pageInfo autoReq.CustomApiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := customApiService.GetCustomApiInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetCustomApiPublic 不需要鉴权的CustomApi接口
// @Tags CustomApi
// @Summary 不需要鉴权的CustomApi接口
// @accept application/json
// @Produce application/json
// @Param data query autoReq.CustomApiSearch true "分页获取CustomApi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customApi/getCustomApiList [get]
func (customApiApi *CustomApiApi) GetCustomApiPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的CustomApi接口信息",
	}, "获取成功", c)
}

func (customApiApi *CustomApiApi) Execute(c *gin.Context) {
	var customApi auto.CustomApi
	err := c.ShouldBindJSON(&customApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	resp, err := customApiService.Execute(&customApi, c.Request.Header.Get("x-token"))
	if err != nil {
		global.GVA_LOG.Error("execute fail", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Result(resp.Code, resp.Data, resp.Msg, c)

}
